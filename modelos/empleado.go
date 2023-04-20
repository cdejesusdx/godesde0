package modelos

type Empleado struct {
	nombre      string
	edad        int
	esMasculino bool
}

func (e *Empleado) Sexo() string {

	return ""
}
