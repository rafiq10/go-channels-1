package spinner

import (
	"fmt"
	"time"
)

func SpinnerMain() {
	const n = 45
	go spinner(1 * time.Second)

	res := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n ", n, res)
}

func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)

}
