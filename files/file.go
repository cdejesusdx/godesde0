package files

import (
	"fmt"
	"os"

	"bufio"
	"io/ioutil"

	"github.com/cdejesusdx/godesde0/ejercicios"
)

var nombreArchivo string = "./files/txt/tabla.txt"

func GrabarTabla() {
	texto := ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumarTabla() {
	texto := ejercicios.TablaDeMultiplicar()
	if !Adjuntar(texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Adjuntar(texto string) bool {
	archivo, err := os.OpenFile(nombreArchivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append")
		return false
	}

	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString")
		return false
	}

	archivo.Close()
	return true
}

func LeerArchivo() {
	archivo, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo))

}

func LeerArchivo2() {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()

}
