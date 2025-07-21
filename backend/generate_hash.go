package main

import (
	"backend/utils"
	"fmt"
)

func main() {
	passwords := []string{"pass456", "long123", "890admin"}

	for _, pw := range passwords {
		hash, err := utils.HashPassword(pw)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Password: %-10s  Hash: %s\n", pw, hash)
	}
}
