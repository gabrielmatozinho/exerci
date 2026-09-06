package main

import (
	"fmt"
)

func classificarTriangulo(A, B, C float64) string {
	if (A+B <= C) || (A+C <= B) || (B+C <= A) {
		return "invalid"
	} else if (A == B) && (B == C) {
		return "equilatero"
	} else if (A == B) || (A == C) || (B == C) {
		return "isosceles"
	} else {
		return "escaleno"
	}
}
func main() {
	var A float64
	var B float64
	var C float64
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Println(classificarTriangulo(A, B, C))
}
