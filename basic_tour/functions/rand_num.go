package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func randomNumGen(length int) string {
	codes := make([]byte, length)

	for i := 0; i < length; i++ {
		// Generate random number 0-9
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			// Fallback or handle error
			return "Error"
		}
		codes[i] = '0' + byte(n.Int64())
	}

	return string(codes)
}

func main() {
	fmt.Println(randomNumGen(6))
}
