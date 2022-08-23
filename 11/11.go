package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	//Объявлнеие множеств
	set1 := []int{1, 2, 3, 4, 5, 6, 10, 12, 8}
	set2 := []int{4, 5, 6, 7, 8, 9, 11, 13, 14, 12, 4, 4}

	m := map[int]int{} //Инициализация мапы
	// if len(set1) > len(set2) {
	// 	set1, set2 = set2, set1
	// }

	//Заполнение мапы. По каждому числу, встречающемуся только один раз, будет значение 1, по другим 2 и больше
	for i := 0; i < len(set1); i++ {
		m[set1[i]]++
	}
	for _, v := range set2 {
		m[v]++
	}
	for key, val := range m {
		if val >= 2 {
			fmt.Println(key)
		}
	}
}
