package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", home)
	puerto := ":3000"
	http.ListenAndServe(puerto, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	archivo := "./index.html"
	http.ServeFile(w, r, archivo)
}

func serverListen(puerto string) {
	fmt.Printf("Server -GO escuchando en el puerto %s", puerto)
}
