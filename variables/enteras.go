package variables

import "fmt"

func MostrarVariablesEnteras() {

	// Variables declarativas:
	var intcomun int

	// Variables por asignacion:
	intde16 := int16(10)
	intde32 := int32(100)
	intde60 := int64(1000)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde16 = ", intde16)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde60 = ", intde60)
}
