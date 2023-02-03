package gameboard

// Liefert die Hauptdiagonale des Spielfelds,
// die von links oben nach rechts unten verläuft.
func (board Board) PrincipalDiag1() Row {
	result := Row{}
	for i, row := range board {
		result = append(result, row[i])
	}
	return result
}

// Liefert die Hauptdiagonale des Spielfelds,
// die von links unten nach rechts oben verläuft.
func (board Board) PrincipalDiag2() Row {
	result := Row{}
	for i := range board {
		result = append(result, board[len(board)-i-1][i])
	}
	return result
}

// Liefert true, falls eine der beiden Hauptdiagonalen nur symbol enthält.
func (board Board) AnyDiagContainsOnly(symbol string) bool {
	return board.PrincipalDiag1().ContainsOnly(symbol) ||
		board.PrincipalDiag2().ContainsOnly(symbol)
}

// Zählt, wie oft symbol in board vorkommt.
func (board Board) Count(symbol string) int {
	result := 0
	for i := range board {
		for _, element := range board[i] {
			if element == symbol {
				result++
			}
		}
	}
	return result
}
