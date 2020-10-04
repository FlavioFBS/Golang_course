package main

import "fmt"

/*Middlewares: interceptores que permiten ejecutar instrucciones comunes a varias funciones
que reciben y devuelve los mismos tipos de variables*/
var result int

func main() {
	fmt.Println("Inicio")
	result = middleware(sumar)(2, 4)
	fmt.Println(result)
	result = middleware(restar)(1, 4)
	fmt.Println(result)
	result = middleware(multiplcar)(2, 4)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplcar(a, b int) int {
	return a * b
}

func middleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Funcion ")
		return f(a, b)
	}
}
