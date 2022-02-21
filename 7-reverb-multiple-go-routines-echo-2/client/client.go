package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	portNum := flag.String("port", "8080", "port to connect to")
	flag.Parse()

	cn, err := net.Dial("tcp", "localhost:"+*portNum)
	if err != nil {
		log.Fatal(err)
	}
	defer cn.Close()
	go mustCopy(os.Stdout, cn)
	mustCopy(cn, os.Stdin)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
