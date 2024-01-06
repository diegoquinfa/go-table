package goTable

import (
	"fmt"
	"strings"
)

const (
	BottomRightCorner = "┘"
	BottomLeftCorner  = "└"

	TopRightCorner = "┐"
	TopLeftCorner  = "┌"

	HorizontalWall = "─"
	VerticalWall   = "│"

	RightIntersection = "┤"
	LeftIntersection  = "├"

	TopIntersection    = "┬"
	MiddleIntersection = "┼"
	BottomIntersection = "┴"
)

type Row struct {
	Columns      []any
	maxLenColumn []int
}

func GetMaxWidthColumns(rows []Row) []int {
	maxWidths := make([]int, len(rows[0].maxLenColumn))

	for _, row := range rows {
		for j, len := range row.maxLenColumn {
			if maxWidths[j] < len {
				maxWidths[j] = len
			}
		}
	}

	return maxWidths
}

func CreateTable(rows []Row) {
	if len(rows) == 0 {
		return
	}

	maxWidths := GetMaxWidthColumns(rows)

	fmt.Print(TopLeftCorner)

	for i, width := range maxWidths {
		for i := 0; i < width+2; i++ {
			fmt.Print(HorizontalWall)
		}
		if i != len(maxWidths)-1 {
			fmt.Print(TopIntersection)
		}
	}

	fmt.Println(TopRightCorner)

	for i, row := range rows {
		if i == 0 {
			for j, column := range row.Columns {
				text := fmt.Sprintf("%v", column)
				len := maxWidths[j] - row.maxLenColumn[j]
				fmt.Printf("%s %s", VerticalWall, text)
				for k := 0; k < len+1; k++ {
					fmt.Print(" ")
				}
			}
			fmt.Println(VerticalWall)

			fmt.Print(LeftIntersection)
			for j, width := range maxWidths {
				for k := 0; k < width+2; k++ {
					fmt.Print(HorizontalWall)
				}
				if j != len(maxWidths)-1 {
					fmt.Print(MiddleIntersection)
				}
			}
			fmt.Println(RightIntersection)
			continue
		}

		for j, column := range row.Columns {
			text := fmt.Sprintf("%v", column)
			spaces := maxWidths[j] - row.maxLenColumn[j]
			fmt.Printf("%s %s", VerticalWall, text)
			for k := 0; k < spaces+1; k++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(VerticalWall)

		if i != len(rows)-1 {
			for j := range row.Columns {
				fmt.Printf("%s", VerticalWall)
				spaces := maxWidths[j]
				for k := 0; k < spaces+2; k++ {
					fmt.Print(" ")
				}
			}

			fmt.Println(VerticalWall)
		}
	}

	fmt.Print(BottomLeftCorner)
	for i, width := range maxWidths {
		for i := 0; i < width+2; i++ {
			fmt.Print(HorizontalWall)
		}
		if i != len(maxWidths)-1 {
			fmt.Print(BottomIntersection)
		}
	}
	println(BottomRightCorner)
}

func NewRow(columns ...any) *Row {
	row := Row{}

	row.Columns = append(row.Columns, columns...)

	for i := range row.Columns {
		text := fmt.Sprintf("%v", columns[i])
		lenText := len([]rune(text))
		if strings.Contains(text, "\x1b[31m") || strings.Contains(text, "\x1b[32m") {
			lenText -= 9
		}
		row.maxLenColumn = append(row.maxLenColumn, lenText)
	}

	return &row
}
