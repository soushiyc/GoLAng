package main

import (
	"fmt"
)

func main(){

	s := "Hello Traveller!"
	fmt.Println(s)

	var protname string

	fmt.Println("Please enter your name: ")
	fmt.Scan(&protname)

	//fmt.Println(protname)

	Prot := Protagonist{protname, weapon[0:1], item[0:1], 1, 10}

	fmt.Println(Prot)
	fmt.Println()

	fmt.Println("Name:", Prot.name)
	fmt.Println("Weapon:", Prot.weapon)
	fmt.Println("Item:", Prot.item)

	fmt.Println("Experience:", Prot.exp)
	fmt.Println("Level:", Prot.level)
	fmt.Println()

	time1 := 0

	for time := 0; time<15 ; time++{
		if Prot.level == 5{
			Prot.weapon = weapon[0:2]
			Prot.item = item[0:2]
			fmt.Println("Name:", Prot.name)
			fmt.Println("Weapon:", Prot.weapon)
			fmt.Println("Item:", Prot.item)
			fmt.Println("Experience:", Prot.exp)
			fmt.Println("Level:", Prot.level)
			fmt.Println()
		}
		if Prot.level == 10{
			Prot.weapon = weapon[0:3]
			Prot.item = item[0:3]
		}
		Prot.exp = Prot.exp*2
		Prot.level = Prot.level + 1
		time1 += time
	}

	fmt.Println("Name:", Prot.name)
	fmt.Println("Weapon:", Prot.weapon)
	fmt.Println("Item:", Prot.item)
	fmt.Println("Experience:", Prot.exp)
	fmt.Println("Level:", Prot.level)
	fmt.Println()

	GoParty := party{"Go Party", Prot}

	fmt.Println(GoParty)
	fmt.Println()
	fmt.Println("Party Name:", GoParty.name)
	fmt.Println("Charactere Name:", GoParty.Character.name)
	fmt.Println("Items:", GoParty.Character.item)
	fmt.Println("Weapons:", GoParty.Character.weapon)
	fmt.Println("Level:", GoParty.Character.level)
	fmt.Println("Experience:", GoParty.Character.exp)
	fmt.Println()

	addressed := addressedParty{"Addressed Party", &Prot}
	fmt.Println(addressed)
	fmt.Println(*addressed.character)
	addressed.character.level = 12
	fmt.Println(addressed.character.level)
	fmt.Println(Prot.level)
	fmt.Println(Prot.exp)
	fmt.Println()

	newRecruit := new(Protagonist)
	newRecruit.name = "Sam"
	newRecruit.item = item[0:1]
	newRecruit.weapon = weapon[0:1]
	newRecruit.level = 1
	newRecruit.exp = 10
	fmt.Println(newRecruit)
	fmt.Println(*newRecruit)
	fmt.Println()

	new_member := newmember{"Lisa", weapon[0:1], item[0:1], 1 , 10}
	new_member.sayHello()
	fmt.Println("Name:", new_member.name)
	fmt.Println("Item:", new_member.item)




}

func (newmember) sayHello()  {
	fmt.Println("Hello Everyone!")
}












/*
1 - 10 10*2
2 - 20
3 - 40
4 - 80
5 - 160

*/

//fmt.Println("Experience:", Prot.exp)
//Prot.level = Prot.exp / (10 * (math.Pow(2, (Prot.level-1))))

//fmt.Println("Level:", Prot.level)
//Prot.exp = Prot

/*var item [3]string
item[0] = "Healing potion"
item[1] = "Magic Potion"
item[2] = "Smoke Bomb"

var weapon [3]string
weapon[0] = "Stick"
weapon[1] = "Dagger"
weapon[2] = "Sword"
*/