package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из
канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	var num int
	c := make(chan int)

	fmt.Print("Введите количество воркеров - ")
	fmt.Scanln(&num)

	go writer(c)
	//Создания пула воркеров
	for i := 0; i < num; i++ {
		go reader(c)
	}

	//Завершение работы по ctrl + c
	ctrl := make(chan os.Signal)
	signal.Notify(ctrl, os.Interrupt, syscall.SIGTERM)
	<-ctrl
	fmt.Println("exit after ctrl c")
	os.Exit(1)
}

//Функция для записи данных в канал
func writer(c chan int) {
	i := 0
	for {
		c <- i
		i++
	}
}

//Функция для чтения и вывода данных из канала
func reader(c chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
