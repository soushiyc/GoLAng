// Golang program to illustrate the
// concept of interfaces
package main

import (
	"fmt"
	"math"
)

func main(){
	var len float64
	var wid float64
	var rad float64
	fmt.Println("Enter the length of the rectangle")
	fmt.Scan(&len)
	fmt.Println("Enter the width of the rectangle")
	fmt.Scan(&wid)
	fmt.Println("Enter the radius of the circle")
	fmt.Scan(&rad)
	c1 := circle{rad}
	r1 := rect{len, wid}

	shapes := []area1{c1, r1}
	per := []perimeter{c1, r1}
	for _, area1 := range shapes {
		fmt.Println(getarea(area1))

	}
	for _, perimeter := range per{
		fmt.Println(getperim(perimeter))
	}



}
type rect struct {
	length float64
	width float64
}

type circle struct {
	radius float64
}

type area1 interface {
	area() float64
}

type perimeter interface {
	perim() float64
}

func (r rect) area() float64{
	fmt.Println("The area of the rectangle is:")
	return r.length * r.width
}

func (c circle) area() float64 {
	fmt.Println("The area of the circle is:")
	return math.Pi * c.radius * c.radius
}

func (rp rect) perim() float64{
	fmt.Println("The perimeter of the rectangle is:")
	return (rp.length + rp.width) * 2
}

func (cp circle) perim() float64 {
	fmt.Println("The perimeter of the circle is:")
	return 2 * math.Pi * cp.radius
}

func getarea(s area1) float64{
	return s.area()
}

func getperim(p perimeter) float64 {
	return  p.perim()
}























/*
// defining an interface
type Sport interface{

	// name of sport method
	sportName() string
}

// declaring a struct
type Human struct{

	// defining struct variables
	name string
	sport string
}

// function to print book details
func (h Human) sportName() string{

	// returning a string value
	return h.name + " plays " + h.sport + "."
}

// main function
func main() {

	// declaring a struct instance
	human1 := Human{"Rahul", "chess"}

	// printing details of human1
	fmt.Println(human1.sportName())

	// declaring another struct instance
	human2 := Human{"Riya", "carrom"}

	// printing details of human2
	fmt.Println(human2.sportName())
}*/

