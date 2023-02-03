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

func ExampleBoard_Count() {
	b1 := Board{
		{"X", "X", "O"},
		{"X", " ", "O"},
		{"X", " ", "O"},
	}

	fmt.Println(b1.Count("X"))
	fmt.Println(b1.Count("O"))
	fmt.Println(b1.Count(" "))
	fmt.Println(b1.Count("U"))

	// Output:
	// 4
	// 3
	// 2
	// 0
}
