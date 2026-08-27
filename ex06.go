package main

import (
	"fmt"
)

func bubbleSort(lista []int) []int {
	n := len(lista)
	for i := 0; i < n; i++ {
		swap := false

		for j := 0; j < n-i-1; j++ {
			if lista[j] > lista[j+1] {
				aux := lista[j]
				lista[j] = lista[j+1]
				lista[j+1] = aux
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
	return lista
}
func main() {
	desordenado := []int{64, 34, 25, 12, 22, 11, 90}
	ordenarCrescente := bubbleSort(desordenado)
	fmt.Println(ordenarCrescente)

}
