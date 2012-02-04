package main

import (
	"fmt"
	"flag"
	"os"
	"net"
	"regexp"
	"encoding/json"

	_ "github.com/davecheney/gmx"
)

var (
	delay = flag.Duration("d", 0, "delay between updates")
	duration = flag.Duration("D", 0, "duration to output continually")
	
	pid = flag.Int("p", 0, "process to inspect")

	socketregex = regexp.MustCompile(`\.gmx\.[0-9]+\.0`)
)

func exitf(values ...interface{}) {
	//fmt.Printf(values...)
	os.Exit(1)
}

func listGmxProcesses() {
	dir, err := os.Open("/tmp")
	if err != nil {
		exitf("unable to open /tmp: %v", err)
	}
	pids, err := dir.Readdirnames(0)
	if err != nil {
		exitf("unable to read pids: %v", err)
	}
	for _, pid := range pids {
		if socketregex.MatchString(pid) {
			c, err := net.Dial("unix", fmt.Sprintf("/tmp/%s", pid))
			if err != nil {
				continue 
			}
			defer c.Close()
			fmt.Fprintln(c, "runtime.version")
			d := json.NewDecoder(c)
			var v string
			if err := d.Decode(&v) ; err != nil {
				continue
			}
			fmt.Printf("%s\t%s\n", pid, v)
		}
	}
}

func main() {
	flag.Parse()
	if *pid == 0 {
		listGmxProcesses()
		return
	}
}
