package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	mapa := map[string]int{"uno": 1, "dos": 2, "tres": 3}
	matriz := [][]int{{1, 2}, {3, 4}, {5, 6}}
	for _, num := range numeros {
		if num%2 == 0 {
			fmt.Printf("Par: %d\n", num)
		} else {
			fmt.Printf("Impar: %d\n", num)
		}
	}
	for clave, valor := range mapa {
		fmt.Printf("Mapa - %s: %d\n", clave, valor)

	}
	for i, fila := range matriz {
		fmt.Printf("Fila %d: %v\n", i, fila)
	}
	fmt.Scanln()
}
