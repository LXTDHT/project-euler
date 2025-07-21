/*
Problem 21:
The sum of all the amicable numbers under 10000
*/

package main

import "fmt"

func main() {
	n := 10000
	sum := 0

	for i := n; i != 0; i-- {
		d := sumOfDivisors(i)
		if sumOfDivisors(d) == i && i != d {
			sum += d + i
		}
	}

	fmt.Println(sum / 2)
}

func sumOfDivisors(n int) int {
	d := 0
	for i := n / 2; i != 0; i-- {
		if n%i == 0 {
			d += i
		}
	}
	return d
}

// Answer: 31626
