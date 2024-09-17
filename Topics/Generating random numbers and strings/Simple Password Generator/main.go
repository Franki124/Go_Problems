package main

import (
	"fmt"
	"math/rand"
)

func main() {
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	upper := upperCase[rand.Intn(len(upperCase))]
	lower := lowerCase[rand.Intn(len(lowerCase))]
	number := numbers[rand.Intn(len(numbers))]
	special := specialSymbol[rand.Intn(len(specialSymbol))]

	password := fmt.Sprintf("%c%c%c%c", upper, lower, number, special)

	fmt.Println(string(password))

}
