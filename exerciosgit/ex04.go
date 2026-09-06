package main

import (
	"errors"
	"fmt"
)

func aplicarDesconto(preco *float64, porcentagem float64) error {
	if porcentagem < 0 || porcentagem > 100 {
		return errors.New("Porcentagem Invalida")
	}

	*preco = *preco - (*preco * porcentagem / 100)

	return nil
}

func main() {
	var x float64 = 500
	var ptr *float64 = &x

	err := aplicarDesconto(ptr, 50)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("%.2f\n", x)
	}
}
