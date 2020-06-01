package mock

import "net"

func (ln *Listener) Accept() (net.Conn, error) {
	ln.AcceptCalled = true

	if ln.AcceptError != nil {
		return nil, ln.AcceptError
	}

	return ln.LN.Accept()
}

func (ln *Listener) Close() error {
	ln.CloseCalled = true

	if ln.CloseError != nil {
		return ln.CloseError
	}

	return ln.LN.Close()
}

func (ln *Listener) Addr() net.Addr {
	ln.AddrCalled = true

	return ln.LN.Addr()
}
