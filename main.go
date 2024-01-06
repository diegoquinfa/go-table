package main

import "fmt"

type table struct {
	Rows []*row
}

type row struct {
	Columns []*column
}

type column struct {
	Content any
}

func main() {

	tb := NewTable(
		NewRow(NewColumn("")),
		NewRow(NewColumn("")),
	)

	NewRow(NewColumn(""))

	fmt.Println(tb)

}

func NewTable(rows ...*row) *table {

	if rows == nil {
		return &table{}
	}

	return &table{
		Rows: rows,
	}
}

func NewRow(columns ...*column) *row {
	return &row{
		Columns: columns,
	}
}

func NewColumn(content any) *column {
	return &column{
		Content: content,
	}
}
