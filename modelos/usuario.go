package modelos

import (
	"time"
)

type Usuario struct {
	Id            int
	Nombre        string
	Estado        bool
	FechaCreacion time.Time
}

func (this *Usuario) AgregarUsuario(id int, nombre string, estado bool, fechaCreacion time.Time) {
	this.Id = id
	this.Nombre = nombre
	this.Estado = estado
	this.FechaCreacion = fechaCreacion
}
