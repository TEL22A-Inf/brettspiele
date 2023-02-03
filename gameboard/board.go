package gameboard

import (
	"fmt"
	"strings"
)

type Board []Row

// Liefert Spalte i.
func (board Board) Column(i int) Row {
	result := Row{}
	for _, row := range board {
		result = append(result, row[i])
	}
	return result
}

// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func (board Board) AnyRowContainsOnly(symbol string) bool {
	for _, row := range board {
		if row.ContainsOnly(symbol) {
			return true
		}
	}
	return false
}

// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func (board Board) AnyColumnContainsOnly(symbol string) bool {
	for i := range board[0] {
		if board.Column(i).ContainsOnly(symbol) {
			return true
		}
	}
	return false
}

func (board Board) String() string {
	width := len(board[0])
	divider := strings.Repeat("+---", width) + "+\n"

	result := divider
	for _, row := range board {
		result += fmt.Sprintf("%s\n", row)
		result += divider
	}

	return result
}
