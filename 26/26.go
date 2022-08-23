package main

import (
	"fmt"
	// "strings"
)

func Unique(str string) bool {
	// str = strings.ToLower(str)
	arr := []byte(str) //Перевод всех символов в байты
	m := make(map[byte]int)

	for _, val := range arr {
		m[val]++ //Если символ не сеществал в мапе, то он запишется как ключ, а его значение станет равно 1.
		//Если существовал, то значение инкрементируется и будет отличным от 1
	}

	//Проверяем мапу на наличие щначений >1
	for key := range m {
		if m[key] > 1 {
			//Если есть, то возвращаем false
			return false
		}
	}

	//Если нет, то возвращаем true
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	str4 := "Aabcd"

	fmt.Println(str1, Unique(str1))
	fmt.Println(str2, Unique(str2))
	fmt.Println(str3, Unique(str3))
	fmt.Println(str4, Unique(str4))
}
