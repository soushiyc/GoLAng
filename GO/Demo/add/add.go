package main

import "fmt"
type addition struct{
	num1 int
	num2 int
	num3 int
}
func main(){

	var numb1 int
	var numb2 int
	var numb3 int
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&numb1)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&numb2)

	addn := addition{numb1, numb2, numb3}
	//fmt.Println(addn)

	type addressedaddn struct {
		name string
		aequation *addition
	}
	addressed := addressedaddn{"Addressed Structure", &addn}
	fmt.Println(addressed)
	fmt.Println(addressed.aequation)

	addressed.aequation.num3 = add(addressed.aequation.num1, addressed.aequation.num2)
	//numb3 = numb1 + numb2

	fmt.Println(addn)
	fmt.Println(addn.num3)

}

func add(x, y int) int{
	return x + y
}