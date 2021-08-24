package main

import "fmt"

func main(){
	var complx1 complex64
	var complx2 complex64
	var complx3 complex64

	fmt.Println("Enter the first complex number")
	fmt.Scan(&complx1)
	fmt.Println(complx1)
	fmt.Println("Enter the second complex number")
	fmt.Scan(&complx2)
	fmt.Println(complx2)
	complx3 = complx1 + complx2
	fmt.Println("Addition:", complx3)
	complx3 = complx1 - complx2
	fmt.Println("Subtraction:", complx3)
	complx3 = complx1 * complx2
	fmt.Println("Multiplication:", complx3)
	complx3 = complx1 / complx2
	fmt.Println("Division:", complx3)
	fmt.Println()
	fmt.Printf("%T, %v.\n", real(complx1), real(complx1))
	fmt.Printf("%T, %v.\n", imag(complx1), imag(complx1))





}
