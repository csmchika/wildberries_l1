package main

import (
	"fmt"
	"strconv"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var num int64
	var i int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)
	fmt.Print("Введите номер бита: ")
	fmt.Scanln(&i)

	if i > 64 || i < 0 {
		panic("Неправильный бит")
	}
	fmt.Println(strconv.FormatInt(num, 2))
	fmt.Println(strconv.FormatInt(1<<(i-1), 2))
	fmt.Println(strconv.FormatInt(num^1, 2))
	num ^= 1 << (i - 1)
	fmt.Println(num)
	fmt.Println(strconv.FormatInt(num, 2))
}
