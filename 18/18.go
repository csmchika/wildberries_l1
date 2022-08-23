package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

type counter struct {
	i int
}

func worker(wg *sync.WaitGroup, m *sync.Mutex, c *counter) {
	m.Lock()   //Блокировка
	c.i++      //Инкрементирование
	m.Unlock() //Разблокировка
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex //Создание мутекса
	c := counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m, &c)
	}

	wg.Wait()

	fmt.Println(c.i)
}
