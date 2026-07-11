package main

import "fmt"

type Sport interface {
	sportName() string
}

type Humano struct {
	name  string
	sport string
}

func (h Humano) sportName() string {
	return h.name + " Plays " + h.sport + "."
}

func main() {
	humano1 := Humano{"Edwar", "Valorant"}

	fmt.Println(humano1.sportName())

	humano2 := Humano{"Carol", "Sprint"}

	fmt.Println(humano2.sportName())
	fmt.Scanln()
}
