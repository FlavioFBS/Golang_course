package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	nombre := "Pepe Perez"
	//agregando 'go' se usa la funcion de modo asincrono
	go miNombreLento(nombre)
	fmt.Println("Escribe algo para finalizar el programa (mientras tanto se irá ejecutando la funcion miNombreLento)")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
