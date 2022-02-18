package main

import (
	"flag"
	"io"
	"log"
	"net"
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
		// handle one connection at a time
		handleConn(cn)
	}
}

// Every second write to net.Conn a string which represents current hour/min/sec
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// io.WriteString accepts io.Writer as a 1st argument
		// net.Conn implements both Writer and Reaer interfaces,
		// so it can be passed as a first argument
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
