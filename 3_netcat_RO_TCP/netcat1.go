package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

// Reads data from the connection to tcp and writes it to the standard output
// until an EOF condicion or an error occurs
// It can be used to test the clock app from the previous section
func main() {
	portNum := flag.String("port", "8080", "port to be dialed to")
	cn, err := net.Dial("tcp", "localhost:"+*portNum)
	if err != nil {
		log.Fatal(err)
	}
	defer cn.Close()
	mustCopy(os.Stdout, cn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// io.Copy copies from src to dst until either EOF is reached on src or an error occurs.
	// It returns the number of bytes copied and the first error encountered while copying, if any.
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
