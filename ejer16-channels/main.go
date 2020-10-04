package main

import (
	"fmt"
	"time"
)

func main() {
	// crear channel
	// make (chan tipoDato)
	// crearemos un canal para calcular el tiempo de ejecucion de un rutina
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegué hasta aquí")

	//msg espera obtener el valor del canal1, el programa no continua hasta que se asigne el valor a msg
	msg := <-canal1

	fmt.Println(msg)
}

func bucle(canal chan time.Duration) {
	//hora de inicio
	inicio := time.Now()
	for i := 0; i < 10000000; i++ {

	}

	//hora de final
	final := time.Now()
	//asignar valor al canal
	canal <- final.Sub(inicio)
}
