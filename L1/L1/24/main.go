package main

import (
	"fmt"
	"math"
)

type Point struct {
	//если название начинается заглавной буквы — это public-доступ,
	//если со строчной — private.
	x, y float64
}

func (p *Point) Get() (x, y float64) {
	return p.x, p.y
}

// конструктор
func (p *Point) Set(x, y float64) {
	p.x = x
	p.y = y
}

// AB = √(xb - xa)2 + (yb - ya)2
func Distance(a, b Point) float64 {
	x1, y1 := a.Get()
	x2, y2 := b.Get()
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func main() {
	var (
		A, B Point
	)
	A.Set(1, 1)
	B.Set(2, 1)
	fmt.Println(Distance(A, B))
}
