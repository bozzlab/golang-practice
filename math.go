package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(1))

	fmt.Println("My favorite number is", rand.Intn(2))

	fmt.Println("My favorite number is", rand.Intn(3))

	fmt.Println("My favorite number is", rand.Intn(4))

	fmt.Println("My favorite number is", rand.Intn(5))

	fmt.Println("My favorite number is", rand.Intn(6))

	fmt.Println("My favorite number is", rand.Intn(7))

	fmt.Println("My favorite number is", rand.Intn(8))

	fmt.Println("My favorite number is", rand.Intn(9))

	rand.Seed(time.Now().UnixNano())

}
