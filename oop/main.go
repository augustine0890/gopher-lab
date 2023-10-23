package main

import (
	"animal/pets"
	"fmt"
	"time"
)

func main() {
	sleepTime := time.Now().Add(time.Duration(-5) * time.Hour)
	pet := pets.NewDog("Oreo", "Black and White", "Mixed", sleepTime)
	fmt.Println(pet.Feed("steak"))
	fmt.Println(pet.GiveAttention("play fetch"))
}
