package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
 а с другой стороны канала — читать. По истечению N секунд программа должна завершаться
*/

func main() {
	var num int
	fmt.Print("Введите количество секунд - ")
	fmt.Scanln(&num)

	c := make(chan int)
	go reader(c)
	go writer(c)

	// time.Sleep(time.Duration(num) * time.Second)
	timer := time.NewTimer(time.Duration(num) * time.Second)
	<-timer.C
	fmt.Println("end")
}

func writer(c chan int) {
	i := 0
	for {
		c <- i
		i++
	}
}

func reader(c chan int) {
	for {
		fmt.Println(<-c)
	}
}
