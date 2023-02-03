package gameboard

import "fmt"

func ExampleBoard_PrincipalDiag1() {
	b1 := Board{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	fmt.Println(b1.PrincipalDiag1())

	// Output:
	// 1 5 9
}

func ExampleBoard_PrincipalDiag2() {
	b1 := Board{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	fmt.Println(b1.PrincipalDiag2())

	// Output:
	// 7 5 3
}
