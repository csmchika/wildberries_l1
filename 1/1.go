package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от
родительской структуры Human (аналог наследования).*/

type Human struct {
	name string
}

func (h Human) SayHello() {
	fmt.Println("I am Human, my name is", h.name)
}

type Action struct {
	Human
}

func main() {
	ac := Action{Human{"Alex"}} // Инициализация структуры
	// ac = Action{Human: Human{"Alex"}}
	// У объекта Action можно неявно вызвать метод принадлежащий Human SayHello()
	ac.SayHello() // I am Human, my name is Alex
	// Так же можно явно вызвать этот метод
	ac.Human.SayHello() // I am Human, my name is Alex
	Human1 := Human{name: "Not Alex"}
	Human1.SayHello() // I am Human, my name is Not Alex
	polic()
}

type SecondAction struct {
	Human
}

func (s SecondAction) SayHello() {
	fmt.Println("I am NOT HUMAN, my name is", s.name)
}
func polic() {
	Human2 := Human{name: "Alex2"}
	SecondA := SecondAction{Human: Human{"Not Alex 2"}}
	Human2.SayHello()        // I am Human, my name is Alex2
	SecondA.SayHello()       // I am NOT HUMAN, my name is Not Alex 2
	SecondA.Human.SayHello() // I am Human, my name is Not Alex 2
}
