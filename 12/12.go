package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var m = make(map[string]struct{})
	var set []string

	//Запись строк в мапу. Существующие слова будут иметь структуру в значении
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = struct{}{}
	}

	//Запись слов в массив
	for key := range m {
		set = append(set, key)
	}

	fmt.Println("Множество -", set)
}
