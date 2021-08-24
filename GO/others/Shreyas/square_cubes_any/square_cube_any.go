package main

import (
	"fmt"
	"math"
)

func main(){
	var input1 int
	var y float64
	var z float64
	var square float64
	var cube float64
	fmt.Println("Please enter a number:")
	fmt.Scan(&input1)
	for input1!=0 {
		y = float64(input1)
		z = math.Mod(y,10)
		square += (z*z)
		cube += (z*z*z)
		input1 = input1 / 10
		fmt.Println(square)
	}
	fmt.Println("The sum of the squares of the inputs is:", square)
	fmt.Println("The sum of the cubes of the inputs is:", cube)
	fmt.Println("The sum of the previous values is:", square+cube)

}