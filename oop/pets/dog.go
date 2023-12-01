package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Name      string
	Color     string
	Breed     string
	lastSlept time.Time
	Animal
}

func (d Dog) needsSleep() bool {
	// return time.Now().Sub(d.lastSlept) > 4*time.Hour
	return time.Since(d.lastSlept).Hours() > 4
}

func (d Dog) sleep() {
	d.lastSlept = time.Now()
}

// func (d Dog) Feed(food string) string {
// return fmt.Sprintf("%s is eating %s", d.Name, food)
// }

func (d Dog) GiveAttention(activity string) string {
	if d.needsSleep() {
		d.sleep()
		return "Your dog is tired and needs to rest"
	}
	response := ""
	switch strings.ToUpper(activity) {
	case "PET":
		response = "wags tail"
	case "Playing Fetch":
		response = "return the ball and jump waiting for you to throw it again"
	default:
		response = "barks"
	}

	return fmt.Sprintf("%s loves attention, %s will cause him to %s", d.Name, response, activity)
}

func NewDog(name, color, breed string, lastSlept time.Time) Pet {
	return Dog{
		Name:      name,
		Color:     color,
		Breed:     breed,
		lastSlept: lastSlept,
		Animal:    Animal{latsAte: time.Now()},
	}
}
