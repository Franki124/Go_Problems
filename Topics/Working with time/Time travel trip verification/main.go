package main

import (
	"fmt"
	"time"
)

func readDate() time.Time {
	// DO NOT MODIFY the code block below!
	var year, month, day int
	fmt.Scan(&year, &month, &day)

	// Write the code here to create a time.Time object from the input
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return date
}

func main() {
	departureDate := readDate()
	historicalEventDate := readDate()
	returnDate := readDate()

	// Write the code here to check if the trip is valid using the `Before()` and `After()` methods:
	if departureDate.After(historicalEventDate) && returnDate.After(historicalEventDate) {
		fmt.Println("Valid trip!")
	} else {
		fmt.Println("Invalid trip!")
	}
}
