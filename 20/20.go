package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func Reverse(s string) string {
	texts := strings.Split(s, " ")
	start, end := 0, len(texts)-1
	for ; start < end; start, end = start+1, end-1 {
		texts[start], texts[end] = texts[end], texts[start]
	}
	return strings.Join(texts, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку:")
	text, _ := reader.ReadString('\n')
	reverse := Reverse(text)
	fmt.Println(reverse)
}
