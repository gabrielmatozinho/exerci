package main

import "fmt"

func buscaLinear(slice []int, alvo int) int {
	for i, valor := range slice {
		if valor == alvo {
			return i
		}
	}
	return -1
}
func main() {
	numeros := []int{14, 5, 89, 2, 7, 50}
	fmt.Println(buscaLinear(numeros, 14))

}
