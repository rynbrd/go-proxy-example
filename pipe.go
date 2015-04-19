package main

import (
	"io"
	"net"
)

const (
	BufSize = 256
)

// Copy data between two connections. Return EOF on connection close.
func Pipe(a, b net.Conn) error {
	done := make(chan error, 1)

	cp := func(r, w net.Conn) {
		var err error
		var n int
		defer func() { done <- err }()
		buf := make([]byte, BufSize)
		for {
			if n, err = r.Read(buf); err != nil {
				break
			}
			if _, err = w.Write(buf[:n]); err != nil {
				break
			}
			logger.Debugf("copied %d bytes from %s to %s", n, r.RemoteAddr(), w.RemoteAddr())
		}
	}

	go cp(a, b)
	go cp(b, a)
	err1 := <-done
	err2 := <-done
	if err1 != io.EOF {
		return err1
	}
	if err2 != io.EOF {
		return err2
	}
	return io.EOF
}
