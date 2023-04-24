package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MostrarNombreEnLento(nombre string, chanel chan bool) {
	letras := strings.Split(nombre, "")
	for i, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("En la posicion %d se encuentra la letra %s \n", i, letra)
	}

	chanel <- true
}
