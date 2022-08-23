package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	var m = map[int]int{}
	var mu = sync.Mutex{}
	wg := sync.WaitGroup{}
	num := 100

	//Создание пула воркеров
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(v int) {
			mu.Lock()
			m[v] = v
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
