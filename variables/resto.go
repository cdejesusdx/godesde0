package variables

import (
	"fmt"
	"strconv"
	"time"
)

// Variables de scope global para otros file que importen el package variables
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Carlos"
	Estado = true
	Sueldo = 85000.78
	Fecha = time.Now()

	fmt.Println("Nombre = ", Nombre)
	fmt.Println("Estado = ", Estado)
	fmt.Println("Sueldo = ", Sueldo)
	fmt.Println("Fecha = ", Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}

func ConvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
