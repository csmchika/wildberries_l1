package main

import (
	"fmt"
	"sync"
	"time"
)

/*
2. Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты
в stdout
*/

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// Простой запуск 5 горутин, вывод рандомный
	for _, n := range arr {
		go fmt.Println(n * n)
	}

	time.Sleep(5 * time.Millisecond)
	fmt.Println("______________")
	SolutionMutex(arr)
	fmt.Println("______________")
	SolutionChannel(arr)

}

func SolutionMutex(nums []int) {
	mut := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, v := range nums {
		go func(n int, m *sync.Mutex) {
			m.Lock()
			fmt.Println(n * n)
			m.Unlock()
			wg.Done()
		}(v, &mut)
	}
	wg.Wait()
}

func SolutionChannel(nums []int) {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c chan<- int, v int) {
			c <- v * v
			wg.Done()
		}(&wg, ch, v)
	}

	// go func(wg *sync.WaitGroup) {
	// 	wg.Wait()
	// 	close(ch)
	// }(&wg)

	// // Чтение из канала результатов пока канал открыт
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	go func(c chan int) {
		for v := range ch {
			fmt.Println(v)
		}
	}(ch)
	wg.Wait()
	close(ch)
}
