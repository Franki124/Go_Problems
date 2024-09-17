package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	factorial := func(number int) int {
		f := 1
		for i := 1; i <= number; i++ {
			f *= i
		}

		return f
	}(number)

	fmt.Println(factorial)
}
