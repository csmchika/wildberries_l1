package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	timer := time.NewTimer(t) //Инициализируем таймер

	<-timer.C //Ждём, когда в канал передастся нужное значение
}

func main() {
	fmt.Println("Start")
	sleep(5 * time.Second) //Передаем в функцию желаемое время сна
	fmt.Println("Stop")
}
