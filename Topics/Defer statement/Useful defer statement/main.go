package main

import "fmt"

func main() {
	var num1, num2, num3 int32

	defer func() {
		if num1 >= 0 && num1 <= 100 {
			fmt.Println("num1: OK")
		} else {
			fmt.Println("num1: ERROR")
		}
		if num2 > 0 {
			fmt.Println("num2: OK")
		} else {
			fmt.Println("num2: ERROR")
		}
		if num3%2 == 0 {
			fmt.Println("num3: OK")
		} else {
			fmt.Println("num3: ERROR")
		}
	}()

	// Do not remove the line below!
	num1, num2, num3 = hiddenValuesHandler()
}
