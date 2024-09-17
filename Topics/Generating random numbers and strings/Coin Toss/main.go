package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed float64
	fmt.Scanln(&seed)
	rand.Seed(int64(seed))
	randomValue := rand.Float64()

	if randomValue <= 0.5 {
		fmt.Println("TAILS")
	} else {
		fmt.Println("HEADS")
	}

}
