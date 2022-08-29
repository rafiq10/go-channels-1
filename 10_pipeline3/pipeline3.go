package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// the call "go counter(naturals)" converts "chan int" into "chan<- int"
	go counter(naturals)
	// the call "go squarer(naturals, squares)" converts "chan int" into "<-chan int" and into "chan<- int" for the second argument
	go squarer(naturals, squares)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}