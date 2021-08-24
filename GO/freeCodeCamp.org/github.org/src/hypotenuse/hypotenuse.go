package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var side1 float64
	var side2 float64
	fmt.Println("Enter the sides:")
	fmt.Scan(&side1)
	fmt.Scan(&side2)
	v := Vertex{side1, side2}
	fmt.Println(v.Abs())
}

