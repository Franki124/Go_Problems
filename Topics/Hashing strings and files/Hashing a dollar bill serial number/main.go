package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strings"
)

func hashSerialNumber(serialNumber string) string {
	// The `count` variable stores the total number of 'E' characters in `serialNumber`:
	count := strings.Count(serialNumber, "E")

	// If the count of 'E' is 0 or odd, then apply SHA256 to `serialNumber` ONCE:
	if count == 0 || count%2 != 0 {
		sha256Hash := sha256.New()
		sha256Hash.Write([]byte(serialNumber))

		serialNumber = fmt.Sprintf("%x", sha256Hash.Sum(nil)) // DO NOT delete!
	}

	// If the count of 'E' is even, then apply MD5 to `serialNumber` iteratively based
	// on the total count of 'E' characters in the `serialNumber`:
	if count%2 == 0 {
		for i := 0; i < count; i++ {
			md5Hash := md5.New()
			md5Hash.Write([]byte(serialNumber))

			serialNumber = fmt.Sprintf("%x", md5Hash.Sum(nil)) // DO NOT delete!
		}
	}
	return serialNumber
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var serialNumber string
	fmt.Scanln(&serialNumber)

	fmt.Println(hashSerialNumber(serialNumber))
}
