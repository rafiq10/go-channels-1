package main

import "fmt"

func main() {
	// Infinite suqarer
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	// Squarer
	go func() {
		for {
			i := <-naturals
			squares <- i * i
		}

	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
