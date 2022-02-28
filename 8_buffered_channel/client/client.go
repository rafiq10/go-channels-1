package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

type closer interface {
	CloseRead() error
}

func main() {
	numPort := flag.String("port", "8080", "port numberto dial")
	flag.Parse()
	cn, err := net.Dial("tcp", "localhost:"+*numPort)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP Session Open")
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, cn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(cn, os.Stdin)
	c, ok := cn.(closer)
	if ok {
		c.CloseRead()
	} else {
		cn.Close()
	}

	<-done // wait for background goroutine to finish
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
