package main

import (
	"animal/pets"
	"fmt"
	"time"
)

func main() {
	sleepTime := time.Now().Add(time.Duration(-3) * time.Hour)
	var animals []pets.Pet
	animals = append(animals, pets.NewCat("Mamba", "White", "Mixed", sleepTime))
	animals = append(animals, pets.NewDog("Oreo", "Black and White", "Mixed", sleepTime))
	for _, pet := range animals {
		if pet.IsHungry() {
			fmt.Println(pet.Feed("steak"))
		} else {
			fmt.Println("Your animal isn't hungry, waiting")
			time.Sleep(3 * time.Second)
			fmt.Println(pet.Feed("kibble"))
		}
		fmt.Println(pet.GiveAttention("play fetch"))
	}
}
