package gameboard

type Row []string

// Liefert true, falls row nur aus dem Zeichen symbol besteht.
func (row Row) ContainsOnly(symbol string) bool {
	for _, character := range row {
		if character != symbol {
			return false
		}
	}
	return true
}
