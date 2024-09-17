package main

import "fmt"

const (
	Min = 0
	Max = 100
)

func main() {
	constraint := newConstraint(Min, Max)

	var number int
	fmt.Scan(&number)

	fmt.Println(constraint(number))
}

func newConstraint(a, b int) func(num int) int {
	minValue := a
	maxValue := b

	return func(num int) int {
		if num < a {
			return minValue
		}
		if num > b {
			return maxValue
		}

		return num
	}
}
