package main

import "fmt"

func main(){
	landmarks := map[string] string{
		"Mondstadt": "Mondstadt Cathedral",
		"Liyue": "Liyue Harbour",
		"Dihua Marsh": "Wangshu Inn",
	}
	if lndmrks, ok := landmarks["Mondstadt"]; ok{
		fmt.Println(lndmrks)
	}
	var loc1 string
	fmt.Println("Please enter the location you would like to search")
	fmt.Scan(&loc1)
	if lndmrks, ok := landmarks[loc1]; ok{
		fmt.Println(lndmrks)
	}else {
		fmt.Println("The requested place is not in the database")
	}

}



