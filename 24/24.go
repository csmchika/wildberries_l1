package main

import (
	"fmt"
	"math"
)

/*
24.	Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором
*/

type Point struct {
	x float64
	y float64
}

//Конструктор
func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}
func main() {
	//Инициализация двух точек
	p1 := NewPoint(1, 1)
	p2 := NewPoint(3, 3)

	//Поиск расстояние через теорему Пифагора
	x := p1.GetX() - p2.GetX()
	y := p1.GetY() - p2.GetY()
	z := math.Sqrt(x*x + y*y)
	fmt.Println("Расстояние -", z)
}
