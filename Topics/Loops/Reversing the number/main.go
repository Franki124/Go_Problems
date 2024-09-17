package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	if input == 0 {
		fmt.Println("0")
		return
	}

	reversedInt := 0
	for input != 0 {
		reversedInt = reversedInt*10 + input%10
		input /= 10
	}

	output := strconv.Itoa(reversedInt)
	fmt.Println(output)
}
