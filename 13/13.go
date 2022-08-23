package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/
func main() {
	a, b := 130, 310
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
	a += b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
