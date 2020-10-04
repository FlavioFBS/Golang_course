package main

import "fmt"

func main() {
	potencia(2)
}

func potencia(numero int) {
	if numero > 10000000 {
		return
	}
	fmt.Println(numero)
	potencia(numero * numero)
}
