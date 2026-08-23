package main

import (
	"fmt"
)

func Buscabinaria(lista []int, alvo int) int {
	inicio := 0
	fim := len(lista) - 1
	for inicio <= fim {
		meio := (inicio + fim) / 2
		if lista[meio] == alvo {
			return meio
		} else if lista[meio] < alvo {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}
func main() {
	lista := []int{3, 7, 12, 19, 25, 31, 42, 55, 68, 74, 83, 91, 105, 120, 135}
	var alvo int

	fmt.Scan(&alvo)

	idx := Buscabinaria(lista, alvo)
	if idx != -1 {
		fmt.Printf("idx: %d | val: %d", idx, lista[idx])
	} else {
		fmt.Println("valor não encontrado")
	}

}
