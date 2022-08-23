package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	Quicksort(ar)
	fmt.Println(ar)
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(arr []int) int {
	pivot := arr[len(arr)/2]

	left := 0
	right := len(arr) - 1

	for {
		for ; arr[left] < pivot; left++ {
		}

		for ; arr[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}
