package main

import "fmt"

func main(){
	var cond bool
	var input int

	fmt.Println("Please enter the order of the result that you would like: ")
	fmt.Println("Enter 'true' for ascending and 'false' for descending")
	fmt.Scan(&cond)
	fmt.Println("Enter the max integer of the series: ")
	fmt.Scan(&input)
	for i := 0; i<=input;i++ {
		if cond == false {
			defer fmt.Println(i)
		}else {
			fmt.Println(i)
		}
	}
}