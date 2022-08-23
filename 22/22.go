package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает,
	делит, складывает, вычитает две числовых переменных a,b,
	значение которых > 2^20.
*/
func main() {
	a := new(big.Float)
	a.SetString("10000000000000000000")
	b := new(big.Float)
	b.SetString("4842981490498425498")

	answer := new(big.Float)

	fmt.Println("Сложение - ", answer.Add(a, b))
	fmt.Println("Умножение - ", answer.Mul(a, b))
	fmt.Println("Вычитание - ", answer.Sub(a, b))
	fmt.Println("Деление - ", answer.Quo(a, b))
}
