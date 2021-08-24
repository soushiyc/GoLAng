package main

import "fmt"

func main(){
	landmarks := map[string] string{
		"Mondstadt": "Mondstadt Cathedral",
		"Liyue": "Liyue Harbour",
		"Dihua Marsh": "Wangshu Inn",
	}
	fmt.Println(landmarks)
	landmarks["Northwest Liyue"] = "Jueyun Karst"
	fmt.Println(landmarks)
	landmarks["Windrise"] = "Falcon Coast"
	fmt.Println(landmarks)
	delete(landmarks, "Windrise")
	fmt.Println(landmarks)


	//	*** m := map[[]int]string{} ***
	//	*** This will show an error as a slice is not a valid key type for mapping. But, we can do it with arrays ***


	m := map[[1]int]string{
		{1}: "One",
		{2}: "Two",
		{3}: "Three",
	}
	fmt.Println()
	fmt.Println(m[[1]int{1}])
	fmt.Println()
	m[[1]int{4}] = "Four"
	fmt.Println(m)
	//fmt.Println(m[[1]int{4}])


	//*** Manipulating a map at one place impacts all other instances of the map
	//so be careful when you're adding/ deleting content from maps ***
	mp := m
	delete(mp, [1]int{4})
	fmt.Println(mp)
	fmt.Println(m)

	fmt.Println()
	_, ok := m[[1]int{4}]
	fmt.Println(ok)
	num, ok := m[[1]int{3}]
	fmt.Println(num, ok)
	fmt.Println(len(m))


}
