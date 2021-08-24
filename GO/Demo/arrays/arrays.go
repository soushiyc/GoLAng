package main

import "fmt"

func main(){
	var a [3] int
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(len(a))
	fmt.Println()

	a[0], a[1] = 1, 2

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println()

	acopy := a
	fmt.Println(a)
	fmt.Println(acopy)
	fmt.Println(len(acopy))
	fmt.Println(acopy == a)
	fmt.Println()

	a[2] = 3
	fmt.Println(a)
	fmt.Println(acopy)
	fmt.Println(len(acopy))
	fmt.Println(acopy == a)
	fmt.Println()

	b := [4]int{1,5,7}
	fmt.Println(b)
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	for _, v := range c {
		fmt.Println(v)
	}
	fmt.Println()

	array1 := [...]int{0: 1}
	fmt.Println(array1)
	array2 := [...]int{2: 2, 4: 6, 7 : 5, 9: 2}
	fmt.Println(array2)

	strarray := [...]string{"Hello", "World", "\b!"}
	fmt.Println(strarray)
	fmt.Println()

	array2d := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(array2d)
	array3d := [3][3][3]int{array2d, array2d, array2d}
	fmt.Println(array3d)
	fmt.Println()


	rray2d := [2][2]int{{1, 2}, {4, 5}}
	fmt.Println(rray2d)
	rray3d := [2][2][2]int{rray2d, rray2d}
	fmt.Println(rray3d)
	fmt.Println()

	newarray := make([]int, 10, 100)
	fmt.Println(len(newarray), cap(newarray))

}
