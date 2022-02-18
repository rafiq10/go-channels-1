package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, host string
}

func (c *clock) watch(w io.Writer, r io.Reader) {
	// if _, err := io.Copy(w, r); err != nil {
	// 	log.Fatal(err)
	// }
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

// Reads data from the connection to tcp and writes it to the standard output
// until an EOF condicion or an error occurs
// It can be used to test the clock app from the previous section
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: netcat2 NAME=PORT ...")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, v := range os.Args[1:] {
		fields := strings.Split(v, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", v)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}
	for _, cl := range clocks {
		cn, err := net.Dial("tcp", "localhost:"+cl.host)
		if err != nil {
			log.Fatal(err)
		}
		defer cn.Close()
		go cl.watch(os.Stdout, cn)
	}
	// sleep while other goroutines work
	for {
		time.Sleep(15 * time.Second)
	}

}
