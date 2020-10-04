package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rutaRelativa := "./archivo.txt"
	leerArchivo2(rutaRelativa)
	archivo, extension := "ejm_guardar", "txt"
	guardarArchivo(archivo, extension)
	guardarArchivo2(archivo, extension)
}

func leerArchivo(rutaRelativa string) {
	archivo, err := ioutil.ReadFile(rutaRelativa)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo")
	} else {
		fmt.Println("Contenido del archivo:")
		fmt.Println(string(archivo))
	}
}

func leerArchivo2(rutaRelativa string) {
	archivo, err := os.Open(rutaRelativa)
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		// revisar linea x linea
		numeroLinea := 1
		fmt.Println("Contenido del archivo:")
		for scanner.Scan() {
			linea := scanner.Text()
			fmt.Printf("Linea %d> %s\n", numeroLinea, linea)
			numeroLinea++
		}
	}
	archivo.Close()
}

func guardarArchivo(nombreArchivo string, extension string) {
	rutaActual := "./"
	archivo, err := os.Create(rutaActual + nombreArchivo + "." + extension)
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	// guardar en el archivo
	fmt.Fprintln(archivo, "Linea texto para guardar")
	archivo.Close()
}

func guardarArchivo2(nombreArchivo string, extension string) {
	rutaActual := "./"
	fileName := rutaActual + nombreArchivo + "." + extension
	contenido := "\n Segunda linea de ejemplo"
	if Append(fileName, contenido) == false {
		fmt.Println("Hubo un error en la segunda linea")
	}

}

//Append ...
func Append(archivo string, texto string) bool {
	// se abrira el archivo en modo solo lectura-escritura, y para que no se limpie y se abra al final se usa O_APPEND
	// 0644: permiso de escritura, lectura
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	// grabar el string dentro del archivo
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
