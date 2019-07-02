// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }

// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// func main() {
// 	fmt.Println(m)
// 	fmt.Printf("%v\n", m["Google"].Lat)
// 	fmt.Printf("%T\n", m["Google"].Lat)
// }

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
