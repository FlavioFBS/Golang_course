package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)
	paises["Peru"] = "Lima"
	paises["Chile"] = "Santiago"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)

	curso := map[string]int{
		"Algoritmica":       15,
		"Algebra":           12,
		"Matematica basica": 16}
	// agregando:
	curso["Calidad"] = 20
	//borrando:
	delete(curso, "Algebra")
	fmt.Println(curso)

	for curso, nota := range curso {
		fmt.Printf("Nota del curso %s: %d\n", curso, nota)
	}
	// consultar existencia:
	nota, existe := curso["Computacion"]
	fmt.Printf("Como el valor de existencia del curso es %t, su nota es %d", existe, nota)
}
