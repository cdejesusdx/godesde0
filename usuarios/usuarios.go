package usuarios

import (
	"fmt"
	"time"

	"github.com/cdejesusdx/godesde0/modelos"
)

func AgregarUsuario() {
	u := new(modelos.Usuario)
	u.AgregarUsuario(1, "Carlos Bautista", true, time.Now())
	fmt.Println(u)
}
