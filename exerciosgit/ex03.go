package main

import (
	"fmt"
)

func analisarIntervalo(inicio, fim int) (pares int, impares int, somaTotal int) {

	for i := inicio; i <= fim; i++ {
		somaTotal += i
		if i%2 == 0 {
			pares++
		} else {
			impares++
		}
	}
	return pares, impares, somaTotal
}

func main() {
	var inicio int
	var fim int

	fmt.Scan(&inicio)
	fmt.Scan(&fim)
	pares, impares, somaTotal := analisarIntervalo(inicio, fim)
	fmt.Println(pares, impares, somaTotal)

}
