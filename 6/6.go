package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины
*/

func main() {
	cancelWithCloseV1()
	cancelWithBool()
	cancelWithContext()
}

func cancelWithCloseV1() {
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("1 return")
				return
			default:
				fmt.Println("1")
			}
		}
	}()

	time.Sleep(1 * time.Microsecond)
	close(quit)
}

func cancelWithBool() {
	die := make(chan bool)

	go func() {
		for {
			select {
			case <-die:
				die <- true
				fmt.Println("2 return")
				return
			default:
				fmt.Println("2")
			}
		}
	}()

	time.Sleep(1 * time.Microsecond)
	die <- true

	// Ждем, пока все горутины закончат выполняться
	<-die
	time.Sleep(1 * time.Microsecond)
}

func cancelWithContext() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 3
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("3 return")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())

	for n := range gen(ctx) {
		fmt.Println(n)
		if n > 6 {
			break
		}
	}
	cancel()
}
