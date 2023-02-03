package gameboard

// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func AnyRowContainsOnly(board []Row, symbol string) bool {
	for i := range board {
		if ContainsOnly(board[i], symbol) {
			return true
		}
	}
	return false
}
