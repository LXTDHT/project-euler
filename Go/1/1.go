// Sum of numbers divisible by 3 or 5 under 1000

package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
