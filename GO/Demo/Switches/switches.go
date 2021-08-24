package main

import (
	"fmt"
)

func main(){
	locations := []string{"Monstad", "Liyue", "Stormbearer Mountains", "Springvale", "Wangshu Inn", "Roaming"}
	fmt.Println(locations)

	for _, location := range locations{
/*		if location == "Monstad"{
			action("Let's go to the library")
		} else if location == "Liyue"{
			action("Let's go meet Ninghuang")
		} else if location == "Stormbearer Mountains"{
			action("Let's lay down of the soft grass")
		} else if location == "Springvale"{
			action("Let's climb up the windmill")
		} else if location == "Wangshu Inn"{
			action("Let's order a feast")
		}
*/
		switch location  {
		case "Monstad":
			action("Let's go to the library")
		case "Liyue":
			action("Let's go meet Ninghuang")
		case "Stormbearer Mountains":
			action("Let's lay down of the soft grass")
		case "Springvale":
			action("Let's climb up the windmill")
		case "Wangshu Inn":
			action("Let's order a feast")
		default:
			action("Let's go defeat some Hilichurls")
		}
	}
}

func action(action string){
	fmt.Println(action)
	fmt.Println()
}
/*func meetingning(meetingning string){
	fmt.Println(meetingning)
	fmt.Println()
}
func relax(relax string){
	fmt.Println(relax)
	fmt.Println()
}
func windmill(windmill string){
	fmt.Println(windmill)
	fmt.Println()
}
func feast(feast string){
	fmt.Println(feast)
	fmt.Println()
}*/