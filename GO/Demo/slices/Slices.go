package main

import "fmt"

func main(){
	fixed := [...]int{1, 2, 3, 4, 5}
	fmt.Println(fixed)
	fmt.Println()
	slices := []int{8, 3, 2, 2, 6, 6, 6, 5, 6, 2}
	fmt.Println(slices)
	fmt.Println(len(slices))
	fmt.Println()
	slices = []int{9, 2, 8, 3, 2, 1}
	fmt.Println(slices)
	fmt.Println(len(slices))
	fmt.Println()
	slices = append(slices, 3, 3, 2, 0)
	fmt.Println(slices)
	fmt.Println(len(slices))
	slices = []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(cap(slices), len(slices))

	slice2 := make([]int, 5)
	slice2 = []int{0,1,2}
	fmt.Println(cap(slice2), len(slice2))
	fmt.Println()

	/*
	Slice of a slice
	*/
	fmt.Println(slices)
	slices = slices[0:7]
	fmt.Println(slices)

}
