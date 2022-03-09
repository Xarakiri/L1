package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	a := big.NewInt(1048577)
	b := big.NewInt(1048577)

	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Div(a, b))
	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Sub(a, b))
}
