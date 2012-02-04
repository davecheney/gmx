package gmx

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
)

func serve(l net.Listener, r *Registry) {
	go func() {
		// if listener is a unix socket, delete it on shutdown
		if l, ok := l.(*net.UnixListener); ok {
			if a, ok := l.Addr().(*net.UnixAddr); ok {
				defer os.Remove(a.Name)
			}
		}
		defer l.Close()
		for {
			c, err := l.Accept()
			if err != nil {
				return
			}
			go handle(c, r)
		}
	}()
}

func handle(c net.Conn, reg *Registry) {
	defer c.Close()
	r := bufio.NewReader(c)
	w := json.NewEncoder(c)
	for {
		v, _, err := r.ReadLine()
		if err != nil {
			return
		}
		if err := w.Encode(reg.value(string(v))()); err != nil {
			// close connection on error
			return
		}
	}
}
