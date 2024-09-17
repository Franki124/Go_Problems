package main

import "fmt"

func main() {
	var start, inc int
	fmt.Scan(&start, &inc)

	increment := newIncrement(start, inc)
	fmt.Println(increment(inc))
}

func newIncrement(start int, inc int) func(inc int) int {
	number := start
	return func(inc int) int {
		number += inc
		return number
	}
}
