package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	portNum := flag.String("port", "8080", "port to listen to")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*portNum)
	log.Printf("listener address: %s", listener.Addr())
	if err != nil {
		log.Fatal(err)
	}

	for {
		cn, err := listener.Accept()
		if err != nil {
			// connection aborted
			log.Print(err)
			continue
		}
		// handle one connections CONCURRENTLY
		go handleConn(cn)
	}
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}
