package main

import "fmt"

func generarCircunferencia[r int | float32](radius r) {
	c := 2 * 3 * radius
	fmt.Println("La circunferencia es de: ", c)
}

func main() {
	var radius1 int = 8
	var radius2 float32 = 9.5

	generarCircunferencia(radius1)
	generarCircunferencia(radius2)
	fmt.Scanln()
}
