package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

type TypeDefiner interface {
}

func getType(td TypeDefiner) string {
	return fmt.Sprintf("Type is %T", td)
}
func main() {

	//Объявление разных типов
	var a int
	var b string
	var c bool
	var d chan string

	spoter(a)
	spoter(b)
	spoter(c)
	spoter(d)

	var typeDef TypeDefiner

	typeDef = 0
	fmt.Println(getType(typeDef))

	typeDef = "abc"
	fmt.Println(getType(typeDef))

	typeDef = true
	fmt.Println(getType(typeDef))

	typeDef = make(chan int)
	fmt.Println(getType(typeDef))
}

func spoter(val interface{}) {
	fmt.Println(reflect.TypeOf(val))
}
