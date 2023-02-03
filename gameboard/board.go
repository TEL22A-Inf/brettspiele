package gameboard

type Board []Row

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
