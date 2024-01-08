package main

import (
	"strings"

	"github.com/diegoquinfa/goTable/goTable"
)

func ColumnSplit(text string) []string {
	splitedText := strings.Split(text, " ")

	content := make([]string, 0)

	var n, i uint16
	for int(i) < len(splitedText) {

		if n == 3 {
			content = append(content, strings.Join(splitedText[i-n:i], " "))
			n = 0
		}

		i++
		n++
	}

	// Agregar las palabras restantes
	if n > 0 {
		content = append(content, strings.Join(splitedText[i-n:], " "))
	}

	return content
}
func main() {

	tb, _ := goTable.NewTable(
		goTable.NewRow("díégo esto es una prueba de un texto enorme ultra mega enorme", "Quintana", "hola", "hola"),
		goTable.NewRow("Dairin", "Fajardo", "hola", "decimo", "mundo acost"),
		goTable.NewRow("marcela", "Cortez", "hola"),
		goTable.NewRow("carmen", "", "Rodigrez", "", "mundo"),
		goTable.NewRow("Jesus", "mundo"),
	)

	tb.Print()
}
