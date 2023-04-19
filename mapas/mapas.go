package mapas

import "fmt"

func MostrarMapas() {
	/* Mapa declarativo:
	paises := make(map[string]string)

	paises["Republica Dominicana"] = "Santo Domingo"
	paises["Argentina"] = "Buenos Aries"

	fmt.Println(paises)
	fmt.Println(paises["Republica Dominicana"])
	*/

	/* Mapa por asignacion: */
	// Liga de Béisbol Profesional de la República Dominicana
	lidon := map[string]int{
		"Tigres del Licey":     0,
		"Águilas Cibaeñas":     10,
		"Estrellas Orientales": 2,
		"Leones del Escogido":  1,
		"Toros del Este":       5,
		"Gigantes del Cibao":   8,
	}

	// fmt.Println(lidon)
	// Iterando el mapa
	/*
		for equipo, carreras := range lidon {
			fmt.Printf("Equipo %s, tiene %d carrera(s)\n", equipo, carreras)
		}
	*/

	// Eliminando elemento del mapa
	delete(lidon, "Tigres del Licey")
	fmt.Println(lidon)

	// Buscando un elemento en el mapa
	carreras, existe := lidon["Águilas Cibaeñas"]
	fmt.Printf("Carrera(s) encontrada(s) %d, y el equipo existe = %t \n", carreras, existe)
}
