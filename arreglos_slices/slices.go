package arreglos_slices

import "fmt"

var tablaSlice []int = []int{5, 10, 15}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MostrarSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   // Slice creado desde un vector, desde la posicion 3 al final
	porcion2 := arreglo[:5]  // Slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // Slice creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Printf("Porcion 1 %d \n", porcion)
	fmt.Printf("Porcion 2 %d \n", porcion2)
	fmt.Printf("Porcion 3 %d \n", porcion3)

}
