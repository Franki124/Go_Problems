package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	// Find the shadowed variable within the anonymous function below and fix it:
	increment := func(number int) {
		result := number * 2 // nolint: gocritic
		result++
	}

	increment(number)
	fmt.Println(number)
}
