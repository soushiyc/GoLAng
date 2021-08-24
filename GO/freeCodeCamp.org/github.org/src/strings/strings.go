package main

import "fmt"

func main(){
	str1 := "Hello"
	str2 := "How are you?"
	byte1 := []byte(str1)
	byte2 := []byte(str2)
	fmt.Println(byte1)
	fmt.Println(byte2)
	rune1 := 'N'
	fmt.Printf("%T,%v.\n", rune1, rune1)		//We get int32 as the type because 'runes' are just a type-alias for int32
}
