package main

import (
	"fmt"
	"sync"
	"time"
)

type task3 struct {
	arr   []int
	mutex sync.Mutex
	wg    sync.WaitGroup
}

/*
3. Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/
func main() {
	arr := []int{2, 4, 6, 8, 10}
	var sum int

	for _, n := range arr {
		go square(&sum, n)
	}

	time.Sleep(5 * time.Millisecond)

	fmt.Println(sum)
	SolutionMutex(arr)
	SolutionChannel(arr)
}

func square(sum *int, n int) {
	*sum += n * n
	// fmt.Println(*sum)
}

func SolutionMutex(nums []int) {
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
	fmt.Println("sum Mutex :=", sum)
}

func SolutionChannel(nums []int) {
	ch := make(chan int, 1)
	task := task3{arr: nums}
	ch <- 0
	for _, value := range task.arr {
		task.wg.Add(1)
		go func(w *sync.WaitGroup, v int, c chan int) {
			c <- (v * v) + <-c
			task.wg.Done()
		}(&task.wg, value, ch)
	}
	task.wg.Wait()
	close(ch)
	fmt.Println("sum Channel :=", <-ch)

}
