package main

import "fmt"

type animalia struct {
	Name string
	Place string
}

/* All Speeds are in MPH */

type mammalia struct{
	kingdom animalia
	canFly bool
	canSwim bool
	egglaying bool
	speed float64
}

type aves struct {
	kingdom animalia
	canFly bool
	canSwim bool
	flyspeed float64
	swimspeed float64
}

type reptilia struct {
	kingdom animalia
	havelegs bool
	isvenomous bool
	canswim bool
	speed float64
}

func main()  {
	/* Reptiles */
	anim_Kingcobra := animalia{"King Cobra", "South Asia"}
	//anim_BlackMumba := animalia{"Black Mamba", "Sub-Saharan Africa"}
	anim_GalapTortoise := animalia{"Galapagos Tortoise", "Galapagos Islands"}
	king_cobra := reptilia{anim_Kingcobra, false, true, true, 12}
	black_mamba := reptilia{animalia{Name: "Black Mamba", Place: "Sub-Saharan Africa"}, false, true, true, 16}
	galap_tortoise := reptilia{anim_GalapTortoise, true,false, false, 0.16}

	fmt.Println(king_cobra.kingdom.Name)
	fmt.Println(black_mamba.kingdom.Name)
	fmt.Println(galap_tortoise.kingdom.Name)
	fmt.Println()

	/* Mammals */
	echidna := mammalia{animalia{Name: "Echidna", Place: "Australia"}, false, false, true, 1.5}
	polar_bears := mammalia{animalia{"Polar Bear", "Nprth Pole"}, false, true, false, 25}
	fmt.Println(echidna)
	fmt.Println(polar_bears)
}
