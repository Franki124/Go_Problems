package main

import (
	"fmt"
	"time"
)

func main() {
	var year, month, day int
	var year2, month2, day2 int
	// Chicken born time:
	fmt.Scan(&year, &month, &day)
	chicken := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	// Egg born time:
	fmt.Scan(&year2, &month2, &day2)
	egg := time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.UTC)

	switch {
	case chicken.Before(egg):
		fmt.Println("Chicken is first!")
	case chicken.After(egg):
		fmt.Println("Egg is first!")
	default:
		fmt.Println("Time collapsed")
	}
}
