package main

import "fmt"

type numbers struct{
	num1 float64
	num2 float64
	num3 float64
}

func main(){
	var input1 float64
	var input2 float64
	var input3 float64

	fmt.Println("Please enter three numbers")
	fmt.Scan(&input1)
	fmt.Scan(&input2)
	fmt.Scan(&input3)

	numstruct := numbers{input1, input2, input3}

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