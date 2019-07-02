package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	//fmt.Printf((board[0]))

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[0][1] = "X"
	board[1][1] = "C"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println(cap(board[0]), len(board[0]))
	fmt.Println(cap(board[1]), len(board[1]))
	fmt.Println(cap(board[2]), len(board[2]))
	fmt.Println(cap(board), len(board))
	// one dimension but use \n for look like new line in terminal
}
