package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hello World!"
	s2 := "Hello All.."
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[11])
	fmt.Printf("%c", s[11])
	fmt.Println()
	fmt.Println(s[0:6])
	fmt.Println(s[6:])
	fmt.Println(s[6:11])

	s1 := s[:11] + " again"
	fmt.Println(s1)

	s1 += "!"
	fmt.Println(s1)
	fmt.Println()
	fmt.Println("Hello \nWorld!")
	fmt.Println("Hello \tWorld!")
	fmt.Println("Hello \bWorld!")
	fmt.Println()
	str := "string"
	by := []byte(str)
	fmt.Println(str, by)
	fmt.Printf("%s, %s", str, by)

	fmt.Println("\n")
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.HasPrefix(s, "Hello"))
	fmt.Println(strings.HasSuffix(s, "!"))
	fmt.Println(strings.HasSuffix(s, "World"))
	fmt.Println(strings.Replace(s, "World", "All!", 1))
	fmt.Println(strings.Replace(s2, ".", "!", 1))
	fmt.Println(strings.Replace(s2, ".", "!", 2))
	fmt.Println()

	var strb strings.Builder
	fmt.Println("Hello", strb.String())
	fmt.Println(strb.Cap())
	fmt.Println(strb.Len())
	strb.WriteString("World!")
	fmt.Println(strb.String())
	fmt.Println("Hello", strb.String())
	fmt.Println(strb.Cap())
	fmt.Println(strb.Len())
	strb.Grow(16)
	fmt.Println(strb.Cap())
	fmt.Println(strb.Len())
	strb.Reset()
	fmt.Println()

	fmt.Println(strb.Cap())
	fmt.Println(strb.Len())
	fmt.Println(strb.String())

	strb.Write([]byte{65, 66, 67})
	fmt.Println(strb.String())
	fmt.Println()

	x := 832266656
	y := strconv.Itoa(x)
	fmt.Printf("%s , %s ", x, y)
	fmt.Println()
	fmt.Printf("%T, %T", x, y)

	a := "8322666562"
	a2 := "83226665"
	b, _ := strconv.Atoi(a)
	fmt.Printf(" %s , %s ", a, b)
	fmt.Println()
	fmt.Printf("%T, %T", a, b)

	fmt.Println(strings.Compare(a2, a))

}
