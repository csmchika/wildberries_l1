package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func Remove(s []int, i int) ([]int, error) {
	if i >= 0 && i < len(s) {
		// создаем новый срез
		ret := make([]int, 0)
		// копируем до i-го элемента
		ret = append(ret, s[:i]...)
		// копируем после i-го элемента и возвращаем
		return append(ret, s[i+1:]...), nil
	} else {
		return []int{}, fmt.Errorf("invalid index")
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i int
	fmt.Print("Введите i: ")
	fmt.Scanln(&i)
	fmt.Println("До - ", s)
	newSlice, ok := Remove(s, i)
	if ok == nil {
		fmt.Println("После - ", newSlice)
	} else {
		fmt.Println(ok)
	}
}
