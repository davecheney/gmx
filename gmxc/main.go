package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"time"
	//	"strings"
)

var (
	delay    = flag.Duration("d", 0, "delay between updates")
	duration = flag.Duration("D", 0, "duration to output continually")

	pid = flag.Int("p", 0, "process to inspect")

	socketregex = regexp.MustCompile(`\.gmx\.[0-9]+\.0`)
)

type conn struct {
	net.Conn
	*json.Decoder
	*json.Encoder
}

func dial(addr string) (*conn, error) {
	c, err := net.Dial("unix", addr)
	return &conn{
		c,
		json.NewDecoder(c),
		json.NewEncoder(c),
	}, err
}

func listGmxProcesses() {
	dir, err := os.Open("/tmp")
	if err != nil {
		log.Fatalf("unable to open /tmp: %v", err)
	}
	pids, err := dir.Readdirnames(0)
	if err != nil {
		log.Fatalf("unable to read pids: %v", err)
	}
	for _, pid := range pids {
		if socketregex.MatchString(pid) {
			c, err := dial(fmt.Sprintf("/tmp/%s", pid))
			if err != nil {
				continue
			}
			defer c.Close()
			c.Encode([]string{"os.args"})
			var result = make(map[string]interface{})
			if err := c.Decode(&result); err != nil {
				log.Printf("unable to decode response from %s: %v", pid, err)
				continue
			}
			if args, ok := result["os.args"]; ok {
				fmt.Printf("%s\t%q\n", pid, args)
			}
		}
	}
}

func main() {
	flag.Parse()
	if *pid == 0 {
		listGmxProcesses()
		return
	}
	c, err := dial(fmt.Sprintf("/tmp/.gmx.%d.0", *pid))
	if err != nil {
		log.Fatalf("unable to connect to process %d: %v", *pid, err)
	}
	defer c.Close()
	deadline := time.Now().Add(*duration)
	for {
		if err := c.Encode(flag.Args()); err != nil {
			log.Fatalf("unable to send request to process: %v", err)
		}
		var result = make(map[string]interface{})
		if err := c.Decode(&result); err != nil {
			log.Fatalf("unable to decode response: %v", err)
		}
		for k, v := range result {
			fmt.Printf("%s: %v\n", k, v)
		}
		if time.Now().After(deadline) {
			return
		}
		time.Sleep(*delay)
	}
}
