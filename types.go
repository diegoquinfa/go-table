package table

type table struct {
	Rows            []*row
	MaxWidthColumns []uint8
}

type row struct {
	Columns    []*column
	NumColumns uint8
	MaxHeight  uint8
	IsHeader   bool
}

type column struct {
	Content []string
	Width   uint8
	Height  uint8
}
