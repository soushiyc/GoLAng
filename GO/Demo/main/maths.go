package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	var in1 int
	var in2 int
	var i bool = true
	fmt.Println(math.Abs(-3))			//3
	fmt.Println(math.Max(3, 14))		//14
	fmt.Println(math.Min(-3, 5))		//-3
	fmt.Println(math.Ceil(3.141))		//4
	fmt.Println(math.Floor(3.845))		//3
	fmt.Println(math.Mod(20, 18))		//2
	fmt.Println(math.Exp(4))				//
	fmt.Println(math.Sqrt(16))			//4
	fmt.Println(math.Pow(5, 3))		//125
	fmt.Println(math.Pow10(5))			//100000
	fmt.Println(math.Mod(15, 24))
	fmt.Println(rand.Intn(15))
	fmt.Println()
	fmt.Println("Enter dividend")
	fmt.Scan(&in1)
	fmt.Println("Enter divisor")
	for i == true{
		fmt.Scan(&in2)
		if in2 == 0{
			fmt.Println("The divisor cannot be Zero.\nPlease enter a non-zero integer")
		} else {
			i = false
		}
	}
	fmt.Println(in1 / in2, in1 % in2,"/",in2)
}