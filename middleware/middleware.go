package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Iniciando...")

	resultado := operacionesMiddleware(sumar)(3, 2)
	fmt.Printf("La suma es: %d \n", resultado)

	resultado = operacionesMiddleware(sumar)(4, 2)
	fmt.Printf("La resta es: %d \n", resultado)

	resultado = operacionesMiddleware(sumar)(5, 4)
	fmt.Printf("La multiplicacion es: %d \n", resultado)
}

func operacionesMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Iniciando Operacion: ")
		return f(a, b)
	}
}
