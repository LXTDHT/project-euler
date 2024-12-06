/*
Problem 5:
Smallest number divisible by all numbers from 1 to 20
*/

package main

import "fmt"

func main() {
	// All numbers divisible by x are divisible by all even/whole (?) dividends of x
	var x int = 20
	for div, n := x, x; div > (x / 2); div-- {
		if n%div != 0 {
			n += x
			div = x
		}
		fmt.Println(n)
	}
}

// Answer: 232792560
