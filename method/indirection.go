package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)        // 6 8
	ScaleFunc(&v, 10) // 60 80

	p := &Vertex{4, 3} // x =4 y =3
	p.Scale(3)         // 12 9
	ScaleFunc(p, 8)    // 96 72

	fmt.Println(v, p)
	//{60 80} &{96 72}

	// p := &Vertex{4, 3} // x =4 y =3
	// p.Scale(3)         // 12 9
	// ScaleFunc(p, 8)    // 96 72

	//{60 80} {96 72}
}
