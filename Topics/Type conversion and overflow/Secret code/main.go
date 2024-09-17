package main

import "fmt"

func main() {
	var s string
	var shift int
	fmt.Scan(&s, &shift)

	for _, ch := range s {
		convertedToInt := int(ch)
		addedShift := convertedToInt + shift
		enigma := string(rune(addedShift))
		fmt.Print(enigma) // add the shift value to the character value
	}
}
