package main

var item [3]string = [3]string{"Healing potion", "Magic Potion", "Smoke Bomb"}
/*
item[0] = "Healing potion"
item[1] = "Magic Potion"
item[2] = "Smoke Bomb"
*/

var weapon [3]string = [3]string{"Stick", "Dagger", "Sword"}
/*
weapon[0] = "Stick"
weapon[1] = "Dagger"
weapon[2] = "Sword"
*/

type Protagonist struct{
	name string
	weapon []string
	item []string
	level int
	exp int
}

/* Character name, Weapon Possessed, Item Possessed, Player level, Experience points*/

type party struct{
	name string
	Character Protagonist
}
/* Party name, Characters in the party*/

type addressedParty struct{
	name string
	character *Protagonist
}
/* Addressed Structure, the address of the 'Character' structure*/

type newmember struct{
	name string
	weapon []string
	item []string
	level int
	exp int
}













/*var item []string
item := make([]string, 3)
//{"Healing potion", "Magic Potion", "Smoke Bomb"}
item[0] = "Healing potion"
item[1] = "Magic Potion"
item[2] = "Smoke Bomb"

weapon := make([]string, 3)
//weapon :=[]string {"Stick", "Dagger", "Sword"}
weapon[0] = "Stick"
weapon[1] = "Dagger"
weapon[2] = "Sword"

*/


/*var item1 []string = "Healing potion"
var item2 []string = "Magic Potion"
var item3 []string = "Smoke Bomb"

var weapon1 []string = "Stick"
var weapon2 []string = "Dagger"
var weapon3 []string = "Sword"
*/