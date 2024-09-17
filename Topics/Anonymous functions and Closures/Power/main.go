package main

import (
	"fmt"
)

func main() {
	var base, exponent int
	fmt.Scan(&base, &exponent)

	pow := func(base, exponent int) int {
		result := 1
		for i := 0; i < exponent; i++ {
			result *= base
		}
		return result
	}

	fmt.Println(pow(base, exponent))
}
