package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*var item [3]string
	item[0] = "Healing potion"
	item[1] = "Magic Potion"
	item[2] = "Smoke Bomb"

	var weapon [3]string
	weapon[0] = "Stick"
	weapon[1] = "Dagger"
	weapon[2] = "Sword"

	s := "Hello World!"
	fmt.Println(s)


	var protname string

	fmt.Println("Please enter your name: ")
	fmt.Scan(&protname)

	fmt.Println(protname)

	Prot := Protagonist{protname, weapon[0:1], item[0:1], 1, 10}

	fmt.Println(Prot)
	fmt.Println()

	fmt.Println("Name:", Prot.name)
	fmt.Println("Weapon:", Prot.weapon)
	fmt.Println("Item:", Prot.item)

	fmt.Println("Experience:", Prot.exp)
	fmt.Println("Level:", Prot.level)
	fmt.Println()
*/

	type character struct {
		name   string
		weapon string
		//item   string
		level  int
		//exp    int
	}

	sFrom := `{"name":"wallace", "weapon":"a", "level":1}`
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
