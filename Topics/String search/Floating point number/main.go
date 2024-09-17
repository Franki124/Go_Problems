package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}

	switch {
	case strings.Contains(input, ","):
		fmt.Println("comma detected")
	case strings.Count(input, ".") > 1:
		fmt.Println("multiple dot symbols")
	default:
		fmt.Println("correct")
	}
}
