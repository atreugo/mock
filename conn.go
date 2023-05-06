package mock

import (
	"net"
	"time"
)

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

func (c *Conn) Close() error {
	return c.ErrClose
}

func (c *Conn) Read(b []byte) (int, error) {
	if c.ErrRead != nil {
		return 0, c.ErrRead
	}

	return c.RBuff.Read(b)
}

func (c *Conn) Write(b []byte) (int, error) {
	if c.ErrWrite != nil {
		return 0, c.ErrWrite
	}

	return c.WBuff.Write(b)
}

func (c *Conn) RemoteAddr() net.Addr {
	return zeroTCPAddr
}

func (c *Conn) LocalAddr() net.Addr {
	return zeroTCPAddr
}

func (c *Conn) SetDeadline(_ time.Time) error {
	return c.ErrSetDeadline
}

func (c *Conn) SetReadDeadline(_ time.Time) error {
	return c.ErrSetReadDeadline
}

func (c *Conn) SetWriteDeadline(_ time.Time) error {
	return c.ErrSetWriteDeadline
}
