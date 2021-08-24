package main

import "fmt"

func main(){
	var n map[string]string
	fmt.Println(n == nil)
	n =	map[string]string{}
	fmt.Println(n == nil)
	fmt.Println(n)
	fmt.Println(len(n))
	var m map[string]string
	m = make (map[string]string, 5)
	fmt.Println(len(m))
	m = map[string]string{"Soushi": "Anemo"}
	fmt.Println(m)
	fmt.Println(len(m))
	m["Amber"] = "Pyro"
	m["Noel"] = "Geo"
	m["Kaeya"] = "Cryo"
	fmt.Println(m)
	fmt.Println()
	delete(m, "Kaeya")
	m["Diona"] = "Cryo"
	fmt.Println(m)
	m["Soushi"] += "+Geo"
	fmt.Println()
	fmt.Println(m["Soushi"])
	fmt.Println(m)
	fmt.Println()

	for name, title := range m{
		fmt.Println(name, title)
	}

	fmt.Println()
	title, ok := m["Kaeya"]
	if ok {
		fmt.Println(title)
	}  else {
		fmt.Println("Couldn't find the requested information")
	}

	fmt.Println()
	m["Kaeya"] = "Cryo"

	title, ok1 := m["Kaeya"]
	if ok1 {
		fmt.Println(title)
	}  else {
		fmt.Println("Couldn't find the requested information")
	}


}
