package main

import (
	"fmt"
	"math"
)

type numbers struct{
	num1 float64
	num2 float64
	num3 float64
}

func main(){
	var input1 float64

	var units float64
	var tens float64
	var hundreds float64

	fmt.Println("Please enter a three digit number")
	fmt.Scan(&input1)

	units = math.Mod(input1, 10)
	tens = (math.Mod(input1,100) - units) / 10
	hundreds = (input1-(math.Mod(input1,100))) / 100
	//fmt.Println(units, tens, hundreds)

	numstruct := numbers{hundreds, tens, units}

	fmt.Println("The sum of the squares of the inputs is:", numstruct.square())
	fmt.Println("The sum of the cubes of the inputs is:", numstruct.cube())
	fmt.Println("The sum of the previous values is:", numstruct.cube()+numstruct.square())
}

func (num numbers) square() float64 {
	return ((num.num1*num.num1) + (num.num2*num.num2) + (num.num3*num.num3))
}
func (num numbers) cube() float64 {
	return ((num.num1*num.num1*num.num1) + (num.num2*num.num2*num.num2) + (num.num3*num.num3*num.num3))
}