package gmx

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
)

func serve(l net.Listener, r *Registry) {
	go func() {
		// if listener is a unix socket, clean it up
		if l, ok := l.(*net.UnixListener); ok {
			if a, ok := l.Addr().(*net.UnixAddr); ok {
				defer os.Remove(a.Name)
			}
		}
		defer l.Close()
		log.Printf("listening on %v", l)
		for {
			c, err := l.Accept()
			if err != nil {
				log.Printf("%v exited with %v", l, err)
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
		if err := w.Encode(reg.getter(string(v))()); err != nil {
			log.Printf("%v, unable to write %v", c, err)
		}
	}
}
