// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xyproto/quic (interfaces: QuicSession)

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	tls "crypto/tls"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/xyproto/quic/internal/protocol"
)

// MockQuicSession is a mock of QuicSession interface
type MockQuicSession struct {
	ctrl     *gomock.Controller
	recorder *MockQuicSessionMockRecorder
}

// MockQuicSessionMockRecorder is the mock recorder for MockQuicSession
type MockQuicSessionMockRecorder struct {
	mock *MockQuicSession
}

// NewMockQuicSession creates a new mock instance
func NewMockQuicSession(ctrl *gomock.Controller) *MockQuicSession {
	mock := &MockQuicSession{ctrl: ctrl}
	mock.recorder = &MockQuicSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuicSession) EXPECT() *MockQuicSessionMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method
func (m *MockQuicSession) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream
func (mr *MockQuicSessionMockRecorder) AcceptStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockQuicSession)(nil).AcceptStream), arg0)
}

// AcceptUniStream mocks base method
func (m *MockQuicSession) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream
func (mr *MockQuicSessionMockRecorder) AcceptUniStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockQuicSession)(nil).AcceptUniStream), arg0)
}

// Close mocks base method
func (m *MockQuicSession) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockQuicSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQuicSession)(nil).Close))
}

// CloseWithError mocks base method
func (m *MockQuicSession) CloseWithError(arg0 protocol.ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError
func (mr *MockQuicSessionMockRecorder) CloseWithError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockQuicSession)(nil).CloseWithError), arg0, arg1)
}

// ConnectionState mocks base method
func (m *MockQuicSession) ConnectionState() tls.ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(tls.ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState
func (mr *MockQuicSessionMockRecorder) ConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockQuicSession)(nil).ConnectionState))
}

// Context mocks base method
func (m *MockQuicSession) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockQuicSessionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQuicSession)(nil).Context))
}

// GetVersion mocks base method
func (m *MockQuicSession) GetVersion() protocol.VersionNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(protocol.VersionNumber)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockQuicSessionMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockQuicSession)(nil).GetVersion))
}

// HandshakeComplete mocks base method
func (m *MockQuicSession) HandshakeComplete() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete
func (mr *MockQuicSessionMockRecorder) HandshakeComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockQuicSession)(nil).HandshakeComplete))
}

// LocalAddr mocks base method
func (m *MockQuicSession) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr
func (mr *MockQuicSessionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockQuicSession)(nil).LocalAddr))
}

// OpenStream mocks base method
func (m *MockQuicSession) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream
func (mr *MockQuicSessionMockRecorder) OpenStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockQuicSession)(nil).OpenStream))
}

// OpenStreamSync mocks base method
func (m *MockQuicSession) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync
func (mr *MockQuicSessionMockRecorder) OpenStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockQuicSession)(nil).OpenStreamSync), arg0)
}

// OpenUniStream mocks base method
func (m *MockQuicSession) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream
func (mr *MockQuicSessionMockRecorder) OpenUniStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockQuicSession)(nil).OpenUniStream))
}

// OpenUniStreamSync mocks base method
func (m *MockQuicSession) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync
func (mr *MockQuicSessionMockRecorder) OpenUniStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockQuicSession)(nil).OpenUniStreamSync), arg0)
}

// RemoteAddr mocks base method
func (m *MockQuicSession) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockQuicSessionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockQuicSession)(nil).RemoteAddr))
}

// closeForRecreating mocks base method
func (m *MockQuicSession) closeForRecreating() protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeForRecreating")
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// closeForRecreating indicates an expected call of closeForRecreating
func (mr *MockQuicSessionMockRecorder) closeForRecreating() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForRecreating", reflect.TypeOf((*MockQuicSession)(nil).closeForRecreating))
}

// closeRemote mocks base method
func (m *MockQuicSession) closeRemote(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeRemote", arg0)
}

// closeRemote indicates an expected call of closeRemote
func (mr *MockQuicSessionMockRecorder) closeRemote(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeRemote", reflect.TypeOf((*MockQuicSession)(nil).closeRemote), arg0)
}

// destroy mocks base method
func (m *MockQuicSession) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy
func (mr *MockQuicSessionMockRecorder) destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockQuicSession)(nil).destroy), arg0)
}

// earlySessionReady mocks base method
func (m *MockQuicSession) earlySessionReady() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "earlySessionReady")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// earlySessionReady indicates an expected call of earlySessionReady
func (mr *MockQuicSessionMockRecorder) earlySessionReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "earlySessionReady", reflect.TypeOf((*MockQuicSession)(nil).earlySessionReady))
}

// getPerspective mocks base method
func (m *MockQuicSession) getPerspective() protocol.Perspective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPerspective")
	ret0, _ := ret[0].(protocol.Perspective)
	return ret0
}

// getPerspective indicates an expected call of getPerspective
func (mr *MockQuicSessionMockRecorder) getPerspective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPerspective", reflect.TypeOf((*MockQuicSession)(nil).getPerspective))
}

// handlePacket mocks base method
func (m *MockQuicSession) handlePacket(arg0 *receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket
func (mr *MockQuicSessionMockRecorder) handlePacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockQuicSession)(nil).handlePacket), arg0)
}

// run mocks base method
func (m *MockQuicSession) run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "run")
	ret0, _ := ret[0].(error)
	return ret0
}

// run indicates an expected call of run
func (mr *MockQuicSessionMockRecorder) run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "run", reflect.TypeOf((*MockQuicSession)(nil).run))
}
