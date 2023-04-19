package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{15, 10, 5} // Arreglo de 10 pisiciones
var matriz [10][20]int                 // Matriz: Arreglos de arreglos

func MostrarArreglos() {
	tabla[7] = 53
	tabla[2] = 100

	tabla2 := [10]string{"Carlos", "Jose", "Maria"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[5][10] = 18
	fmt.Println(matriz)
}
