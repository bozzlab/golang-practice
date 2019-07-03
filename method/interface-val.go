package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	var k I

	i = &T{"Hello"}
	describe(i)
	i.M()

	k = F(math.Pi)
	describe(k)
	k.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
