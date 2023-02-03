package gameboard

import "fmt"

func ExampleBoard_AnyRowContainsOnly() {
	board := Board{
		{" ", "X", " "},
		{"X", "X", "X"},
		{" ", " ", " "},
	}

	fmt.Println(board.AnyRowContainsOnly("X"))
	fmt.Println(board.AnyRowContainsOnly("O"))

	// Output:
	// true
	// false
}

func ExampleBoard_AnyColumnContainsOnly() {
	board := Board{
		{" ", "X", " "},
		{"X", "X", " "},
		{" ", "X", " "},
	}

	fmt.Println(board.AnyColumnContainsOnly("X"))
	fmt.Println(board.AnyColumnContainsOnly("O"))

	// Output:
	// true
	// false
}
