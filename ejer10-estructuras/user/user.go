package user

import "time"

//Usuario ...
type Usuario struct {
	ID        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//user hará referencia a la structura Usuario

//AltaUsuario .
func (user *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	user.ID = id
	user.Nombre = nombre
	user.FechaAlta = fechaAlta
	user.Status = status
}
