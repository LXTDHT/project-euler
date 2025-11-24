/*
Problem 6:
The difference between the sum of the squares of the first one hundred natural numbers (1-100) and the square of the sum
*/

package main

import "fmt"

func main() {
	var sum, sq int
	for i := 1; i <= 100; i++ {
		sum += i * i
		sq += i
	}
	fmt.Println(sq*sq - sum)
}

// Answer: 25164150
