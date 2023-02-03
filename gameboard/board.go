package gameboard

type Board []Row

// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func AnyRowContainsOnly(board Board, symbol string) bool {
	for _, row := range board {
		if ContainsOnly(row, symbol) {
			return true
		}
	}
	return false
}
