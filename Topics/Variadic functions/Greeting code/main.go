package main

import "fmt"

func main() {
	var prefix, name1, name2 string
	fmt.Scan(&prefix, &name1, &name2)

	for _, line := range Greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func Greeting(prefix string, name ...string) []string {
	var result []string

	for _, s := range name {
		result = append(result, prefix+" "+s)
	}

	return result
}
