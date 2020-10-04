package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	fmt.Println()
	numero, estado, message := dos(2)
	fmt.Printf("%d --- %t --- %s", numero, estado, message)
	fmt.Println("-------------------------")
	fmt.Println(calculo(1, 2, 3, 4, 5))
	fmt.Println(calculo(10, 45, 1, 5, 16, 89, 2))
}

func uno(numero int) int {
	return numero * 2
}

// retorna lista de tipos de dato
func dos(numero int) (int, bool, string) {
	if numero == 1 {
		return 5, false, "Failed"
	}
	return 10, true, "Success"
}

func calculo(numero ...int) int {
	total := 0
	// range recorre por cada elemento
	for _, num := range numero {
		total += num
	}
	return total
}
