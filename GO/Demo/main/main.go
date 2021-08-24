package main

import (
	"fmt"
)

var three = 3


func main1() {
	var i int = 13
	var p *int = &i
	type celsius int
	type fahrenheit int
	var c celsius = 5
	var f fahrenheit
	fmt.Println("Hello World!")
	fmt.Println(*p)
	fmt.Println(p)
	f = fahrenheit((c * 9 / 5) + 32)
	fmt.Println(f)
	fmt.Println(three)

}

