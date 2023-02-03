package gameboard

import "fmt"

func ExampleAnyRowContainsOnly() {
	b1 := [][]string{
		{" ", "X", " "},
		{"X", "X", "X"},
		{" ", "X", " "},
	}

	fmt.Println(AnyRowContainsOnly(b1, "X"))
	fmt.Println(AnyRowContainsOnly(b1, "O"))

	// Output:
	// true
	// false
}