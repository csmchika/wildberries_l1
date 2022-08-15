package main

import (
	"fmt"
	"sync"
	"time"
)

type task3 struct {
	arr   [5]int
	mutex sync.Mutex
	wg    sync.WaitGroup
}

/*
3. Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/
func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var sum int

	for _, n := range arr {
		go square(&sum, n)
	}

	time.Sleep(5 * time.Second)

	fmt.Println(sum)
	SolutionMutex(arr)
}

func square(sum *int, n int) {
	*sum += n * n
	fmt.Println(*sum)
}

func SolutionMutex(nums [5]int) {
	var sum int
	task := task3{arr: nums}
	task.wg.Add(len(task.arr))
	for _, v := range task.arr {
		go func(n int) {
			task.mutex.Lock()
			sum += n * n
			task.mutex.Unlock()
			task.wg.Done()
		}(v)
	}
	task.wg.Wait()
	fmt.Println("sum :=", sum)
}
