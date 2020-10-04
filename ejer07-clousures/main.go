package main

import "fmt"

// usando una variable de tipo funcion, su valor será una funcion anonima
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5+7 = %d \n", calculo(5, 7))

	//redifiniendo la funcion:
	calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Printf("Sumo 5*7 = %d \n", calculo(5, 7))

	//aplicando closure
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}
}

func operaciones() {
	resultado := func() int {
		a, b := 13, 12
		return a + b
	}
	fmt.Println(resultado())
}

//closures:
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
