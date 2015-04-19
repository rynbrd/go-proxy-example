package main

import (
	"net"
	"runtime"
	"sync/atomic"
)

const (
	BufSize = 256
)

// Copy data between two connections. Return EOF on connection close.
func Pipe(a, b net.Conn) error {
	done := make(chan error)
	var stop int32
	defer func() {
		atomic.StoreInt32(&stop, 1)
	}()

	cp := func(r, w net.Conn) {
		var err error
		var n int
		buf := make([]byte, BufSize)
		for {
			if atomic.LoadInt32(&stop) == 1 {
				return
			}

			if n, err = r.Read(buf); err != nil {
				done <- err
				return
			}
			if _, err = w.Write(buf[:n]); err != nil {
				done <- err
				return
			}
			logger.Debugf("copied %d bytes from %s to %s", n, r.RemoteAddr(), w.RemoteAddr())
			runtime.Gosched()
		}
	}

	go cp(a, b)
	go cp(b, a)
	return <-done
}
