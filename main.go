package main

import (
	"fmt"

	"github.com/cdejesusdx/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertToText(455)

	fmt.Println(estado)
	fmt.Println(texto)
}
