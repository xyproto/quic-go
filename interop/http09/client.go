package http09

import (
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/idna"

	"github.com/xyproto/quic"
)

// RoundTripper performs HTTP/0.9 roundtrips over QUIC.
type RoundTripper struct {
	mutex sync.Mutex

	TLSClientConfig *tls.Config
	QuicConfig      *quic.Config

	clients map[string]*client
}

var _ http.RoundTripper = &RoundTripper{}

// RoundTrip performs a HTTP/0.9 request.
// It only supports GET requests.
func (r *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodGet {
		return nil, errors.New("only GET requests supported")
	}

	r.mutex.Lock()
	hostname := authorityAddr("https", hostnameFromRequest(req))
	if r.clients == nil {
		r.clients = make(map[string]*client)
	}
	c, ok := r.clients[hostname]
	if !ok {
		tlsConf := r.TLSClientConfig.Clone()
		tlsConf.NextProtos = []string{h09alpn}
		c = &client{
			hostname: hostname,
			tlsConf:  tlsConf,
			quicConf: r.QuicConfig,
		}
		r.clients[hostname] = c
	}
	r.mutex.Unlock()
	return c.RoundTrip(req)
}

// Close closes the roundtripper.
func (r *RoundTripper) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, c := range r.clients {
		if err := c.Close(); err != nil {
			return err
		}
	}
	return nil
}

type client struct {
	hostname string
	tlsConf  *tls.Config
	quicConf *quic.Config

	once    sync.Once
	sess    quic.Session
	dialErr error
}

func (c *client) RoundTrip(req *http.Request) (*http.Response, error) {
	c.once.Do(func() {
		c.sess, c.dialErr = quic.DialAddr(c.hostname, c.tlsConf, c.quicConf)
	})
	if c.dialErr != nil {
		return nil, c.dialErr
	}
	return c.doRequest(req)
}

func (c *client) doRequest(req *http.Request) (*http.Response, error) {
	str, err := c.sess.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	cmd := "GET " + req.URL.Path + "\r\n"
	if _, err := str.Write([]byte(cmd)); err != nil {
		return nil, err
	}
	if err := str.Close(); err != nil {
		return nil, err
	}
	rsp := &http.Response{
		Proto:      "HTTP/0.9",
		ProtoMajor: 0,
		ProtoMinor: 9,
		Request:    req,
		Body:       ioutil.NopCloser(str),
	}
	return rsp, nil
}

func (c *client) Close() error {
	if c.sess == nil {
		return nil
	}
	return c.sess.Close()
}

func hostnameFromRequest(req *http.Request) string {
	if req.URL != nil {
		return req.URL.Host
	}
	return ""
}

// authorityAddr returns a given authority (a host/IP, or host:port / ip:port)
// and returns a host:port. The port 443 is added if needed.
func authorityAddr(scheme string, authority string) (addr string) {
	host, port, err := net.SplitHostPort(authority)
	if err != nil { // authority didn't have a port
		port = "443"
		if scheme == "http" {
			port = "80"
		}
		host = authority
	}
	if a, err := idna.ToASCII(host); err == nil {
		host = a
	}
	// IPv6 address literal, without a port:
	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		return host + ":" + port
	}
	return net.JoinHostPort(host, port)
}
