package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		archivo := "./prueba.txt"
		f, err := os.Open(archivo)

		// se ejecuta antes de salir de la funcion
		defer f.Close()

		if err != nil {
			fmt.Println("error abriendo el archivo")
			os.Exit(1)
		}
	*/

	fmt.Println(",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,")
	ejemploPanic()
}

func ejemploPanic() {
	//sin defer no se podra controlar el programa luego del panic
	defer func() {
		// recover: trae resultado del ultimo panic
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic: \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
