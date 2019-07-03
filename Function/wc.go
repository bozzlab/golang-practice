package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fmt.Println("s input----", s)
	words := strings.Fields(s)
	fmt.Println("input to string field", words)
	m := make(map[string]int)
	fmt.Println("map before loop", m)
	for _, word := range words {
		m[word] += 1
		fmt.Println("word in loop", word)
		fmt.Println("map in count", m[word])
	}
	fmt.Println("map after loop", m)
	return m
}

func main() {
	wc.Test(WordCount)
}
