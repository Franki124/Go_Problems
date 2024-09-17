package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string

	fmt.Scan(&source)

	zeros := strings.Count(source, "0")
	units := strings.Count(source, "1")

	if zeros > units {
		fmt.Println("The count of zeros is bigger")
	} else if units > zeros {
		fmt.Println("The count of units is bigger")
	} else {
		fmt.Println("Counts are equal")
	}
}
