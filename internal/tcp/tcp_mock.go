package tcp

import (
	"bytes"
	"net"
	"time"
)

// MockConn is a mock implementation of net.Conn for testing purposes
type MockConn struct {
	readBuffer  *bytes.Buffer
	writeBuffer *bytes.Buffer
}

func NewMockConn(input string) *MockConn {
	return &MockConn{
		readBuffer:  bytes.NewBufferString(input),
		writeBuffer: new(bytes.Buffer),
	}
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return m.readBuffer.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.writeBuffer.Write(b)
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return nil
}

func (m *MockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) GetWrittenData() string {
	return m.writeBuffer.String()
}
