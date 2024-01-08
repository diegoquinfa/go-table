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
		MaxWidthColumns: getMaxWidthColumns(rows),
	}, nil
}

func (t *table) Print() {
	if len(t.Rows) == 0 {
		return
	}

	for rowIndex, row := range t.Rows {
		t.printRow(row, &rowIndex)
	}

	t.intersection(-1)
}

func (t *table) printRow(row *row, rowIndex *int) {
	t.intersection(int8(*rowIndex))

	for lineIndex := uint8(0); lineIndex < row.MaxHeight; lineIndex++ {
		for colIndex, col := range row.Columns {

			content := t.getLineContent(col, row, &lineIndex)

			leftover := int(t.MaxWidthColumns[colIndex]) - utf8.RuneCountInString(content)
			fill := strings.Repeat(" ", leftover)
			fmt.Printf("%s %s%s ", VerticalWall, content, fill)
		}

		t.fillMissingSpaces(row)

		fmt.Println(VerticalWall)
	}
}

func (t *table) fillMissingSpaces(row *row) {
	diff := len(t.MaxWidthColumns) - len(row.Columns)
	if diff != 0 {
		for i := len(t.MaxWidthColumns) - diff; i < len(t.MaxWidthColumns); i++ {
			fmt.Printf("%s %s ", VerticalWall, strings.Repeat(" ", int(t.MaxWidthColumns[i])))
		}
	}
}

func (t *table) getLineContent(col *column, row *row, rowIndex *uint8) string {
	if *rowIndex == 0 {
		return col.Content[0]
	} else if len(col.Content) == int(row.MaxHeight) {
		return col.Content[*rowIndex]
	}
	return ""
}

func (t *table) intersection(position int8) {
	left := TopLeftCorner
	middle := TopIntersection
	right := TopRightCorner

	if position > 0 {
		left = LeftIntersection
		middle = MiddleIntersection
		right = RightIntersection
	} else if position == -1 {
		left = BottomLeftCorner
		middle = BottomIntersection
		right = BottomRightCorner
	}

	fmt.Print(left)
	for i, width := range t.MaxWidthColumns {
		for j := uint8(0); j < width+2; j++ {
			fmt.Print(HorizontalWall)
		}
		if i != len(t.MaxWidthColumns)-1 {
			fmt.Print(middle)
		}
	}
	fmt.Println(right)
}
