package table

import "testing"

func TestFunction(t *testing.T) {
	table, _ := NewTable(
		NewRow("Primer nombre", "Primer apellido", "Segundo apellido", "Estado Civil", "Descripción"),
		NewRow("Maria", "Gomez", "Perez", "Divorciada", "Mujer trabajadora y dedicada, le gusta la lectura y el arte."),
		NewRow("Juan", "Meléndez", "Puello", "Casado", "Es un hombre amigable con una personalidad extrovertida y una risa contagiosa."),
		NewRow("Carlos", "Salcedo", "", "Soltero", "Ingeniero de Sistemas apasionado por el desarrollo web."),
		NewRow("Ana", "Rodríguez", "Fernández", "Casada", "Madre de dos hijos, disfruta cocinar y hacer ejercicio."),
		NewRow("Luis", "Vargas", "Castro", "Soltero", "Estudiante de medicina con una fuerte inclinación hacia la investigación biomédica."),
		NewRow("Elena", "Martínez", "Santos", "Viuda", "Ama de casa jubilada, disfruta de la jardinería y la lectura."),
		NewRow("Roberto", "Ramírez", "González", "Comprometido", "Viajero empedernido, ha visitado más de 20 países."),
	)

	table.Print()
}
