package gmx

import (
	"fmt"
	_log "log"
	"net"
	"os"
)

const GMX_VERSION = 0

var log = _log.New(os.Stderr, "gmx: ", 0)

var registry = newRegistry()

func init() {
	s, err := localSocket()
	if err != nil {
		log.Printf("unable to open local socket: %v", err)
		return
	}

	// register the registries keys for discovery
	Register("registry", func() interface{} {
		return registry.keys()
	})
	serve(s, registry)
}

func localSocket() (net.Listener, error) {
	return net.ListenUnix("unix", localSocketAddr())
}

func localSocketAddr() *net.UnixAddr {
	return &net.UnixAddr {
		fmt.Sprintf("/tmp/.gmx.%d.%d", os.Getpid(), GMX_VERSION),
		"unix",
	}
}

// Register assocaites the name with the function f. f may
// be recovered by its name by gmx clients.
func Register(name string, f func() interface{}) {
	registry.register(name, f)
}
