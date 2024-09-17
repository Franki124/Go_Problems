package main

import (
	"fmt"
	"math"
)

func main() {
	var source uint64
	_, err := fmt.Scan(&source)
	if err != nil {
		fmt.Println("unsupported type")
		return
	}

	// check if the source value can be converted to int8 value without overflow
	if source <= math.MaxUint8 {
		fmt.Println("uint8")
	}
	if source <= math.MaxUint16 {
		fmt.Println("uint16")
	}
	if source <= math.MaxUint32 {
		fmt.Println("uint32")
	}
	fmt.Println("uint64")

	// add the same check for 16, 32, 64 bits...

	// check if the source value is out of the range for all types

}
