package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	x := s[:0]
	printSlice(x)

	// Extend its length.
	y := s[:4]
	printSlice(y)

	// Drop its first two values.
	z := s[2:]
	printSlice(z)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
