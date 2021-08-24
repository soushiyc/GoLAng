package main

import "fmt"

func main() {
	type q1 struct
	{
		a int
		b int
		c int
		d int
		e int
		f int
	}
	type r1 struct {
		a1 bool
		b1 bool
		c1 bool
		d1 bool
		e1 bool
		f1 bool
	}
	q := q1{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := r1{true, false, true, true, false, true}
	fmt.Println(r)

	type s struct{
		i q1
		b r1
	}
	s1 := s{q, r}
	fmt.Println(s1)
}

