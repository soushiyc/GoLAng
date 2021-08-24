package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	type character struct {
		name   string `json:"fullname"`
		weapon string `json:"weaponpossed""`
		level  int	  `json:"playerlevel"`
	}

	sFrom := `{"fullname":"wallace", "weaponpossed":"Sword", "playerlevel":1}`
	fmt.Println(sFrom)
	fmt.Println()

	var wallace character
	fmt.Println(wallace)
	fmt.Println()

	err := json.Unmarshal([]byte(sFrom), &wallace)
	if err != nil {
		fmt.Println("Shocked Pickachu Face!")
		fmt.Println(err)
	}
	fmt.Println(wallace)


}

