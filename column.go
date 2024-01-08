package table

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func newColumn(text string) *column {
	content := columnSplit(text, 4)
	width := maxWidthColumn(content)
	height := len(content)

	return &column{
		Content: content,
		Width:   width,
		Height:  uint8(height),
	}
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

func getMaxWidthColumns(rows []*row) []uint8 {

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
