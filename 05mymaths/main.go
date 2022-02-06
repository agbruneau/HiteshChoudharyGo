package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// random from math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
