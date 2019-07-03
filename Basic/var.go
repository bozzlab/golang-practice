// package main

// import "fmt"

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

package main

import "fmt"

var i, j = 1, 2

func main() {
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
