package gmx

import (
	"fmt"
	"net"
	"os"
)

type Config struct {
}

// LocalSocket returns a net.Conn connected to the 
// local unix domain socket for the running gmx process
func (c Config) LocalSocket() (net.Listener, error) {
	path := fmt.Sprintf("/tmp/.gmx.%d.0", os.Getpid())
	return net.ListenUnix("unix", &net.UnixAddr{
		path,
		"unix",
	})
}
