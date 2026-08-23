package main

import (
	"fmt"
)

func conversorEFormatadorDeTemperatura(C float64) (float64, float64) {
	var F float64
	var K float64
	F = (C * 1.8) + 32
	K = C + 273.15
	return F, K
}
func main() {
	var C float64
	var F float64
	var K float64
	fmt.Scan(&C)
	F, K = conversorEFormatadorDeTemperatura(C)
	fmt.Printf("%.2f | %.2f\n", F, K)

}
