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
	arr := [5]int{2, 4, 6, 8, 10}
	// Простой запуск 5 горутин, вывод рандомный
	for _, n := range arr {
		go fmt.Println(n * n)
	}

	time.Sleep(5 * time.Second)

	SolutionMutexWG(arr) // решение с Mutex и WaitGroup

}

func SolutionMutexWG(nums [5]int) {
	mut := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, v := range nums {
		go func(n int, m *sync.Mutex) {
			// Использую mutex для того
			// что бы была реальная конкуренция,
			// без mutex при печати результата
			// может случиться такое что несколько потоков
			// в один момент времени захотят вывести к примеру:
			// "123\n" и "456\n" в stdout, но вместо этого
			// может напечатать "1425\n3\n" c одновременным вызовом fmt.Println()
			// в то же время мы этим добились конкуренции.
			m.Lock()
			fmt.Println(n * n)
			m.Unlock()
			wg.Done()
		}(v, &mut)
	}
	wg.Wait()
}
