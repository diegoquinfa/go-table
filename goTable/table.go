package goTable

import (
	"fmt"
	"strings"
)

func (t *table) Print() {
	if len(t.Rows) == 0 {
		return
	}

	for rowIndex, row := range t.Rows {
		t.intersection(int8(rowIndex))

		// for colIndex, col := range row.Columns {
		// 	content := col.Content[0]
		// 	leftover := int(t.MaxWidthColumns[colIndex]) - len([]rune(content))
		// 	fill := strings.Repeat(" ", leftover)
		// 	fmt.Printf("%s %s%s ", VerticalWall, col.Content[0], fill)
		// }

		for i := uint8(0); i < row.MaxHeight; i++ {
			for colIndex, col := range row.Columns {

				var content string

				if i == 0 {
					content = col.Content[0]
				} else if len(col.Content) == int(row.MaxHeight) {
					content = col.Content[i]
				}
				leftover := int(t.MaxWidthColumns[colIndex]) - len([]rune(content))
				fill := strings.Repeat(" ", leftover)
				fmt.Printf("%s %s%s ", VerticalWall, content, fill)
			}

			diff := len(t.MaxWidthColumns) - len(row.Columns)
			if diff != 0 {
				for i := len(t.MaxWidthColumns) - diff; i < len(t.MaxWidthColumns); i++ {
					fmt.Printf("%s %s ", VerticalWall, strings.Repeat(" ", int(t.MaxWidthColumns[i])))
				}
			}

			fmt.Println(VerticalWall)
		}
	}

	t.intersection(-1)
}

func (t *table) intersection(position int8) {
	leftIntersection := TopLeftCorner
	middleIntersection := TopIntersection
	rightIntersection := TopRightCorner

	if position > 0 {
		leftIntersection = LeftIntersection
		middleIntersection = MiddleIntersection
		rightIntersection = RightIntersection
	} else if position == -1 {
		leftIntersection = BottomLeftCorner
		middleIntersection = BottomIntersection
		rightIntersection = BottomRightCorner
	}

	fmt.Print(leftIntersection)
	for i, width := range t.MaxWidthColumns {
		for j := uint8(0); j < width+2; j++ {
			fmt.Print(HorizontalWall)
		}
		if i != len(t.MaxWidthColumns)-1 {
			fmt.Print(middleIntersection)
		}
	}
	fmt.Println(rightIntersection)

}
