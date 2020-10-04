package main

import (
	"fmt"
	"time"

	usRef "./user"
)

// pepe hereda de la structura Usuario del paquete user
type pepe struct {
	usRef.Usuario
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "Flavio FBS", time.Now(), false)
	fmt.Println(user.Usuario)
}
