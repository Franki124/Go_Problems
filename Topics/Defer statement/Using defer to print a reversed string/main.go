package main

import "fmt"

func main() {
	// delete or add new defer statements below!
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Your name is: %s\n", name)
	fmt.Printf("Reversed name: ")

	reverse(name)
}

func reverse(name string) {
	for _, char := range name {
		defer fmt.Printf("%c", char)
	}
}
