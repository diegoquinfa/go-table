package main

import "fmt"

type table struct {
	Rows []*row
}

type row struct {
	Columns any
}

func main() {

	tb := NewTable(
		NewRow("", ""),
		NewRow(""),
	)

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

func NewRow(columns ...any) *row {

	return &row{
		Columns: columns,
	}
}
