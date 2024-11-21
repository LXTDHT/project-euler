// Sum of even Fibonacci under 4 million

package main

import "fmt"

func main() {
	fib := make(map[int]int)
	fib[0] = 1
	fib[1] = 2
	for i := 2; i > 0; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		if fib[i] > 4000000 {
			break
		}
	}
	var sum int
	for i := range fib {
		if fib[i]%2 == 0 {
			sum += fib[i]
		}
	}
	fmt.Println(sum)
}
