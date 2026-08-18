package main

import (
	"fmt"
)

func main() {
	var L1 float64
	var L2 float64
	var L3 float64

	fmt.Scan(&L1)
	fmt.Scan(&L2)
	fmt.Scan(&L3)
	if (L1+L2 > L3) && (L2+L3 > L1) && (L1+L3 > L2) {
		fmt.Printf()

	}
}
