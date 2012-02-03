package gmx

import (
	"fmt"
        _log "log"
	"net"
        "os"
)

var log = _log.New(os.Stderr, "gmx: ", 0)

var registry = newRegistry()

func init() {
	s, err := localSocket()
	if err != nil {
		log.Printf("unable to open local socket: %v", err)
		return
	}

	// register this registry for discovery
	Register("registry", func() interface{} {
		return registry.keys()
	})
	serve(s, registry)
}

// localSocket returns a net.Conn connected to the 
// local unix domain socket for the running gmx process
func localSocket() (net.Listener, error) {
        path := fmt.Sprintf("/tmp/.gmx.%d.0", os.Getpid())
        return net.ListenUnix("unix", &net.UnixAddr{
                path,
                "unix",
        })
}

// Register assocaites the name with the function f. f may
// be recovered by its name by gmx clients.
func Register(name string, f func() interface{}) {
	registry.register(name, f)
}
