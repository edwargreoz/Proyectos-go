package main

import (
	"fmt"
)

type person struct {
	nombre   string
	Apellido string
	edad     int32
}

func (p person) detallePersona() {
	fmt.Printf("El nombre de la persona es %s y su apellido %s ", p.nombre, p.Apellido)
	fmt.Printf("\nEdad: %d\n  ", p.edad)
}

func main() {
	p1 := person{
		"Edwar Agustin",
		"Gonzalez Requejo",
		12}

	p1.detallePersona()

	p1.nombre = "Agustin"
	p1.edad = 21

	p1.detallePersona()

	fmt.Scanln()
}
