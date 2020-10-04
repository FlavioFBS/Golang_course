package main

import "fmt"

var tabla1 [10]int
var matriz1 [2][3]int

func main() {
	tabla1[0] = 1
	tabla1[5] = 30

	tabla2 := [10]int{5, 2, 0, 5, 6, 8}

	fmt.Println(tabla1)
	fmt.Println(tabla2)

	// slice: arreglo dinamico
	//matriz2 := []int{2, 4, 5}
	variante3()
	/*
		for i := 0; i < len(tabla2); i++ {
			fmt.Println(tabla2[i])
		}
	*/
}

// aplicando slices
func variante() {
	lista := [5]int{1, 2, 3, 4, 5}
	porcion1 := lista[3:]  // desde posicion 3 hasta el fin
	porcion2 := lista[:3]  // desde posicion inicio hasta 3
	porcion3 := lista[2:4] // desde posicion 2 hasta 4
	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func variante2() {
	// crear un vector de 5, con espacio reservado(capacidad) de 20
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante3() {
	/** la capacidad es la potencia de 2 más cercana mayor al tamaño del arreglo
	 */
	elementos := make([]int, 0, 0)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}
