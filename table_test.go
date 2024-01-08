package table

import "testing"

func TestFunction(t *testing.T) {
	table, _ := NewTable(
		NewRow("Primer nombre", "Primer apellido", "Segundo apellido", "Descripción"),
		NewRow("Maria", "Gomez", "Perez", "Mujer trabajadora y dedicada, le gusta la lectura y el arte."),
		NewRow("Juan", "Meléndez", "Puello", "Es un hombre amigable con una personalidad extrovertida y una risa contagiosa."),
		NewRow("Carlos", "Salcedo", "", "Ingeniero de Sistemas apasionado por el desarrollo web."),
		NewRow("Ana", "Rodríguez", "Fernández", "Madre de dos hijos, disfruta cocinar y hacer ejercicio."),
		NewRow("Luis", "Vargas", "Castro", "Estudiante de medicina con una fuerte inclinación hacia la investigación biomédica."),
		NewRow("Elena", "Martínez", "Santos", "Ama de casa jubilada, disfruta de la jardinería y la lectura."),
		NewRow("Roberto", "Ramírez", "González", "Viajero empedernido, ha visitado más de 20 países."),
	)

	table.Print()
}
