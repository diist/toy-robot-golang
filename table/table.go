package table

const MAX_X, MAX_Y int = 4, 4

type Table struct {
	MaxX, MaxY int
}

func NewTable() *Table {
	return &Table{MAX_X, MAX_Y}
}
