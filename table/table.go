package table

const MAX_X, MAX_Y int = 4, 4

type Table struct {
	MaxX, MaxY int
}

func NewTable() *Table {
	return &Table{MAX_X, MAX_Y}
}

func IsValidPosition(x int, y int) bool {
	table := NewTable()
	if x > table.MaxX || x < 0 || y > table.MaxY || y < 0 {
		return false
	}
	return true
}
