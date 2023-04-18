package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var err error
var number int
var resultadoFinal string

func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ingrese un numero para la multiplicacion del mismo: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		resultadoFinal += fmt.Sprintf("%d x %d = %d \n", number, i, i*number)
	}

	return resultadoFinal
}
