package main

import (
	"fmt"
	"math"
)


func main(){
	const num int = 8322666562
	fmt.Println(num)
	var numv int = 123456789
	fmt.Println(numv)
	numv =  987654321
	fmt.Println(numv)

	fmt.Println(math.Pi)

	const(
		a = 832
		b = 266
		c = 65
		d = 62
	)
	fmt.Println(a, b, c, d)

	const(
		q = 83226
		w
		e = 66562
		r
	)
	fmt.Println(q, w, e, r)

	const(
		Zero = iota
		one
		two
		three
		four
	)
	fmt.Println(Zero, one, two, three, four)

	const(
		Zero1 = iota + math.E + math.Pi
		one1
		two1
		three1
		four1
	)
	fmt.Println(Zero1, one1, two1, three1, four1)

	const(
		na = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb

	)

	fmt.Println(kb, mb, gb, tb,pb)
	fmt.Println(zb / eb)
	filesize := 5000000000.
	fmt.Printf("%.2fGB", filesize/gb)

}

