package main

import (
	"io"
	"net"
)

// Copy data between two connections. Return EOF on connection close.
func Pipe(a, b net.Conn) error {
	done := make(chan error, 1)

	cp := func(r, w net.Conn) {
		n, err := io.Copy(r, w)
		logger.Debugf("copied %d bytes from %s to %s", n, r.RemoteAddr(), w.RemoteAddr())
		done <- err
	}

	go cp(a, b)
	go cp(b, a)
	err1 := <-done
	err2 := <-done
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}
