package main

import "fmt"

func main() {
	// Infinite suqarer
	naturals := make(chan int, 100)
	squares := make(chan int, 100)

	// Counter
	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		defer close(naturals)
	}()

	// Squarer
	go func() {
		// for {
		// i, ok := <-naturals
		// if !ok {
		// 	break //channel was closed and drained
		// }
		// squares <- i * i
		// }
		for i := range naturals {
			squares <- i * i
		}
		defer close(squares)
	}()

	// Printer (in main goroutine)
	for i := range squares {
		fmt.Println(i)
	}
}
