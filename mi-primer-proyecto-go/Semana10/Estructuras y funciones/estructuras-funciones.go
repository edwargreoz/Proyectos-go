package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

func (p *Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s", p.Nombre)
}

func main() {
	p := &Persona{
		Nombre: "Juan",
		Edad:   30,
		Email:  "juan@ejemplo.com"}
	fmt.Println(p.Saludar())
	fmt.Scanln()
}
