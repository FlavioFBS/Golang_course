package main

import "fmt"

var estado bool

func main() {
	estado = true
	if estado = false; estado == true {
		fmt.Println("estado en true")
	} else {
		fmt.Println("estado en false")
	}

	switch numero := 5; numero {
	case 1:
		fmt.Println("es uno")
	case 2:
		fmt.Println("es dos")
	case 5:
		fmt.Println("es cinco")
	default:
		fmt.Println("defecto")
	}
}
