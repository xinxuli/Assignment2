package main

import (
	"backend/utils"
	"fmt"
)

func main() {
	passwords := []string{"pass123", "bobsecure", "admin456"}

	for _, pw := range passwords {
		hash, err := utils.HashPassword(pw)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Password: %-10s  Hash: %s\n", pw, hash)
	}
}
