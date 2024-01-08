package table

func NewRow(columnsText ...string) *row {
	cols := make([]*column, 0, len(columnsText))

	for _, text := range columnsText {
		cols = append(cols, newColumn(text))
	}

	return &row{
		Columns:    cols,
		NumColumns: uint8(len(cols)),
		MaxHeight:  getMaxHeighRow(cols),
	}
}

func getMaxHeighRow(columns []*column) uint8 {
	var maxHeight uint8

	for _, col := range columns {
		if col.Height > maxHeight {
			maxHeight = col.Height
		}
	}

	return maxHeight
}
