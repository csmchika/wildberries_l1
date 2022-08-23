package main

import (
	"fmt"
	"strings"
)

/*
	К каким негативным последствиям может привести
	данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func mainExample() {
	someFunc()
}
*/
/*
1) Глобальные переменные
2) Мы будем использовать только первые 100 элементов/битов из 1024, оставшиеся будут висеть в памяти и удалятся лишь после выхода justString из области видимости (а дальше см п.1)
*/

func main() {
	var justString string

	v := createHugeString(1 << 10)
	// Решение: клонировать
	justString = strings.Clone(v[:100])
	//justString = v[:100]
	fmt.Println(justString)
}

func createHugeString(size int) string {
	return strings.Repeat(")", size)
}
