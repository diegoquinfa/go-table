package goTable

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func NewTable(rows ...*row) (*table, error) {
	if rows == nil {
		return nil, fmt.Errorf("rows are empty")
	}

	return &table{
		Rows:            rows,
		MaxWidthColumns: GetMaxWidthColumns(rows),
	}, nil
}

func GetMaxWidthColumns(rows []*row) []uint8 {

	var maxNumColumns uint8

	for _, row := range rows {
		if row.NumColumns > maxNumColumns {
			maxNumColumns = row.NumColumns
		}
	}

	maxWidths := make([]uint8, maxNumColumns)
	for _, row := range rows {
		for j, col := range row.Columns {
			if col.Width > maxWidths[j] {
				maxWidths[j] = col.Width
			}
		}
	}

	fmt.Println(maxWidths)

	return maxWidths
}

func NewRow(columnsText ...string) *row {
	cols := make([]*column, 0, len(columnsText))

	for _, text := range columnsText {
		cols = append(cols, NewColumn(text))
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

func NewColumn(text string) *column {
	content := columnSplit(text, 4)
	width := maxWidthColumn(content)
	height := len(content)

	return &column{
		Content: content,
		Width:   width,
		Height:  uint8(height),
	}
}

func columnSplit(text string, maxSize uint8) []string {
	words := strings.Split(text, " ")

	var content []string
	var lineContent []string

	for i, word := range words {
		lineContent = append(lineContent, word)

		if len(lineContent) == int(maxSize) || i == len(words)-1 {
			content = append(content, strings.Join(lineContent, " "))
			lineContent = []string{}
		}
	}

	return content
}

func maxWidthColumn(content []string) uint8 {
	var maxWidth int

	for _, line := range content {
		width := utf8.RuneCountInString(line)
		if width > maxWidth {
			maxWidth = width
		}
	}

	return uint8(maxWidth)
}
