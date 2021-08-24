/*
A program that takes two integers and takes the modulus of them and tells us if the second number is a factor of the first number
** You have to use type conversion
** The divisor cannot be zero
*/
package main

import (
	"fmt"
	"math"
)

func main(){
	var resultcheck bool
	var input1 int
	var input2 int
	var numb1 float64
	var numb2 float64
	var result float64
	var i bool = true

	fmt.Println("Please enter the dividend")
	fmt.Scan(&input1)
	fmt.Println("Please enter the divisor")

	/*
	For loop protocol:

	for initializer; test/conditional; action(ex incrementer) {}
	or you can use
	for test/conditional {}
	or
	for {} if you'd like an infinite loop or there is a break function inside the loop once certain conditions are met
	You can use '_' to say that you are  not using a part of the conditions
	example:
	for _; i<5; i++{}
	or
	for i:=0;_;i++{}
	*/
	for i == true{
		fmt.Scan(&input2)
		if input2 == 0{
			fmt.Println("The divisor cannot be Zero.\nPlease enter a non-zero integer")
		} else {
			i = false
		}
	}
	numb1 = float64(input1)
	numb2 = float64(input2)
	result = math.Mod(numb1, numb2)
	resultcheck = result == 0
	if resultcheck == true {
		fmt.Println("divisor is a factor of the dividend")
	} else {
		fmt.Println("divisor is not a factor of the dividend")
	}
	fmt.Println()

	/* Boolean operations &, |, ^, &^ */

	fmt.Println(input1 & input2)
	fmt.Println(input1 | input2)
	fmt.Println(input1 ^ input2)
	fmt.Println(input1 &^ input2)

	/* Bit shifting*/
	var bitshift int
	fmt.Println()
	fmt.Println("Enter the number to be bit-shifted")
	fmt.Scan(&bitshift)
	fmt.Println(bitshift << 3)
	fmt.Println(bitshift >> 4)


}

/*


func main(){
	var resultcheck bool
	var input1 int
	var input2 int
	var numb1 float64
	var numb2 float64
	var result float64

	for {
		fmt.Println("Please enter the dividend")
		fmt.Scan(&input1)
		input1 := reflect.TypeOf(input1).Kind()
		if input1 == reflect.Int {
			fmt.Println("Please enter the divisor")
			fmt.Scan(&input2)
			input2 := reflect.TypeOf(input2).Kind()
			if input1 == reflect.Int {
				numb1 = float64(input1)
				numb2 = float64(input2)

				result = math.Mod(numb1, numb2)

				resultcheck = result != 0

				if resultcheck == false {
					fmt.Println("divisor is a factor of the dividend")
					break
				} else {
					fmt.Println("divisor is not a factor of the dividend")
					break
				}
			} else {
				fmt.Println("Please enter an integer")
			}
		}else {
			fmt.Println("Please enter an integer")
		}
	}


}*/


/*
case:
input 1 = 10 -- 4'b1010
input 2 = 7  -- 4'b0111

input1 & input2 = 4'b0010 = 2'd02
input1 | input2 = 4'b1111 = 2'd15
input1 ^ input2 = 4'b1101 = 2'd13
input1 &^ input2 = 4'b1000 = 2'd08
*/