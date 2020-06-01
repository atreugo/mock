package mock

import (
	"bytes"
	"net"
)

type Conn struct {
	RBuff bytes.Buffer
	WBuff bytes.Buffer

	ErrClose            error
	ErrRead             error
	ErrWrite            error
	ErrSetDeadline      error
	ErrSetReadDeadline  error
	ErrSetWriteDeadline error

	net.Conn
}

type Listener struct {
	Listener net.Listener

	AcceptError error
	CloseError  error

	AcceptCalled bool
	CloseCalled  bool
	AddrCalled   bool
}
