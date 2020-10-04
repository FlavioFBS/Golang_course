package main

import "fmt"

func main() {
	fmt.Println("\n\nIteracion 1:")
	i := 10
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("\n\nIteracion 2:")
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("\n\nIteracion 3:")
	numero := 0
	for {
		fmt.Println("Continuar")
		fmt.Println("Ingrese numero clave:")
		fmt.Scanln(&numero)
		if numero == 100 {
			fmt.Println("Numero clave correcto...")
			break
		}
	}

	fmt.Println("\n\nIteracion 4:")
	var k = 0
	for k < 15 {
		fmt.Printf("\n Valor de k: %d", k)
		if k == 5 {
			fmt.Printf("se multiplica por 2 \n")
			k = k * 2
			continue
		}
		fmt.Printf("--------Pasó ...\n")
		k++
	}

	fmt.Println("\n\nIteracion 5:") // aplicacion similar de goto
	var l int = 0
RUTINA:
	for l < 10 {
		if l == 4 {
			l = l + 2
			fmt.Println("voy a RUTINA sumando 2 a l")
			goto RUTINA
		}
		fmt.Printf("Valor de l: %d\n", l)
		l++
	}
}
