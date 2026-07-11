package main

import "fmt"

func main() {
	//declaracion explicita
	var activo bool = true
	var disponible bool = false

	// inferencia de tipo
	var validado = true //go infiere que es bool

	//sin valor definido
	var configurado bool //false por defecto

	fmt.Printf("activo: %t, disponible: %t, validado: %t, configurado: %t\n",
		activo, disponible, validado, configurado)

	//operaciones logicas
	resultado := activo && disponible
	negacion := !activo

	fmt.Printf("resultado: %t, negacion %t\n", resultado, negacion)
	fmt.Scanln()
}
