package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

func GeneratePassword(length int, includeUpper, includeNumbers, includeSymbols bool) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("password length must be greater than 0")
	}
	charpool := lowercase
	if includeUpper {
		charpool += uppercase
	}
	if includeNumbers {
		charpool += numbers
	}
	if includeSymbols {
		charpool += symbols
	}

	if len(charpool) == 0 {
		return "", fmt.Errorf("character pool cannot be empty")
	}
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charpool))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}
		password[i] = charpool[index.Int64()]
	}

	return string(password), nil

}

func main()  {
	length := 12
	includeUpper := true
	includeNumbers := true
	includeSymbols := true

	password, err := GeneratePassword(length, includeUpper, includeNumbers, includeSymbols)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Generate password", password)

}
