package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num1 int = 83
	fmt.Printf("%v, %T\n", num1, num1)
	fmt.Println()
	var num2 float32 = 45
	fmt.Printf("%v, %T\n", num2, num2)
	fmt.Println()
	var num3 float64
	num3 = float64(num1)
	fmt.Printf("%v, %T\n", num3, num3)
	fmt.Println()
	var str  string
	str = string(num1)
	fmt.Printf("%v, %T\n", str, str)
	fmt.Println()
	str = strconv.Itoa(num1)
	fmt.Printf("%v, %T", str, str)
}
