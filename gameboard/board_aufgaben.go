package gameboard

// Liefert die Hauptdiagonale des Spielfelds,
// die von links oben nach rechts unten verläuft.
func (board Board) PrincipalDiag1() Row {
	result := Row{}
	// TODO
	return result
}

// Liefert die Hauptdiagonale des Spielfelds,
// die von links unten nach rechts oben verläuft.
func (board Board) PrincipalDiag2() Row {
	result := Row{}
	// TODO
	return result
}

// Liefert true, falls eine der beiden Hauptdiagonalen nur symbol enthält.
func (board Board) AnyDiagContainsOnly(symbol string) bool {
	// TODO
	return true
}

func (board Board) Count(symbol string) int {
	// TODO
	return 0
}
