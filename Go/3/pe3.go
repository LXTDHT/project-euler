/*
Problem 3:
Largest prime factor of 600851475143
*/

// TODO: Find more efficient algorithm to validate large prime numbers.

package main

import "fmt"

const (
	// P int64 = 50
	P   int64 = 600851475143 // max
	MIN int64 = 300425707570 // minimum prime to start from, used in the loop below
)

func main() {
	var x int64 = MIN
	var primes = []int64{}

	// halving P (guaranteed largest factor) for fewer iterations
	for i := 0; x < P/2+1; i++ {
		z := false
		if x > 2 && x%2 == 0 {
			x++
			continue
		}
		for j := int64(3); j < x; j++ {
			if x%j == 0 {
				z = true
				break
			}
		}
		if z {
			x++
			continue
		}
		primes = append(primes, x)
		// primes[i] = x
		x++
		// break
	}

	for _, p := range primes {
		fmt.Println(p)
	}
}

// Answer:
