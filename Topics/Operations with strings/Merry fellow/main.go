package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Which function from the strings package removes spaces?
func trimmer(word string) string {
	return strings.TrimSpace(word)
}

// DO NOT delete or modify the contents of the main function!
func main() {
	dictionary := make([]string, 0)
	var word string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word = scanner.Text()
		word = trimmer(word)
		dictionary = append(dictionary, word)
	}

	for _, word = range dictionary {
		fmt.Println(word)
	}
}
