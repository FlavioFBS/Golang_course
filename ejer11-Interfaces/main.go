package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

/*Especie Humano*/
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	// herencia
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "hombre"
	}
	return "mujer"
}
func (h *hombre) estaVivo() bool { return h.vivo }

/*
func (m *mujer) respirar()    { m.respirando = true }
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "mujer" }
*/

// HumanosRespirando ...
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando\n", hu.sexo())
}

/* Especie Animal*/
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) esCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

//AnimalesRespirar ...
func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

//AnimalesCarnivoros ...
func AnimalesCarnivoros(an animal) int {
	if an.esCarnivoro() {
		return 1
	}
	return 0
}

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	caso4()
}

func caso1() {
	//en este caso se hardcodeo la struc mujer
	Pedro := new(hombre)
	/* la funcion HumanosRespirando recibe una interface, pero como 'hombre'
	implemento la interface 'humano', es correcto
	*/
	fmt.Println("Genero humano")
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)
}

func caso2() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	/* la funcion HumanosRespirando recibe una interface, pero como 'hombre'
	implemento la interface 'humano', es correcto
	*/
	fmt.Println("Especie humano")
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)
}

func caso3() {
	fmt.Println("Especie animal")
	totalCarnivoros := 0
	Firulais := new(perro)
	Firulais.carnivoro = true
	AnimalesRespirar(Firulais)
	totalCarnivoros += AnimalesCarnivoros(Firulais)

	fmt.Printf("Total Carnivoros %d", totalCarnivoros)
}

func caso4() {
	fmt.Println("Especie animal")
	totalCarnivoros := 0
	Firulais := new(perro)
	Firulais.carnivoro = true
	Firulais.vivo = true
	AnimalesRespirar(Firulais)
	totalCarnivoros += AnimalesCarnivoros(Firulais)

	fmt.Printf("Total Carnivoros %d", totalCarnivoros)
	fmt.Printf("\n Estoy vivo = %t", estoyVivo(Firulais))
}
