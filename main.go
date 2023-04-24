package main

import "github.com/cdejesusdx/godesde0/middleware"

func main() {
	/*
		estado, texto := variables.ConvertToText(455)
		fmt.Println(estado)
		fmt.Println(texto)
	*/

	// Obtener el sistema operativo
	// os := runtime.GOOS

	/* Condicionales:
	   if os := runtime.GOOS; os == "Linux" || os == "OS X" {
	   	fmt.Println("Esto no es Windows, es", os)
	   } else {
	   	fmt.Println("Esto es", os)
	   }

	   switch os := runtime.GOOS; os {
	   case "linux":
	   	fmt.Println("Esto es Linux")
	   case "darwin":
	   	fmt.Println("Esto es Darwin")
	   default:
	   	fmt.Printf("%s \n", os)

	   }


	number, message := ejercicios.ConvertToNumber("M")
	println("El numero convertido es", number)
	println("El resultado de la operacion:", message)

	teclado.IngresoNumeros()

	fmt.Printf(ejercicios.TablaDeMultiplicar())

	// files.GrabarTabla()
	// files.SumarTabla()
	//files.LeerArchivo2()
	//funciones.Exponencia(2)

	//arreglos_slices.MostrarSlice()
	//mapas.MostrarMapas()
	//usuarios.AgregarUsuario()

	var estado bool = false
	chanel := make(chan bool)
	go goroutines.MostrarNombreEnLento("Carlos", chanel)
	defer func() {
		estado = <-chanel
	}()

	fmt.Printf("El estado del canal es %t \n ", estado)

	//var x string
	//fmt.Scan(&x)

	webserver.MiWebServer2()*/

	middleware.MiMiddleware()
}
