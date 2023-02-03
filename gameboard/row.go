package gameboard

type Row []string

func ContainsOnly(row Row, symbol string) bool {
	for _, character := range row {
		if character != symbol {
			return false
		}
	}
	return true
}
