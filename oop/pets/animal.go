package pets

import (
	"fmt"
	"time"
)

type Animal struct {
	latsAte time.Time
}

func (a Animal) Feed(food string) string {
	a.latsAte = time.Now()
	return fmt.Sprintf("The animal is eating %s", food)
}

func (a Animal) IsHungry() bool {
	return time.Since(a.latsAte).Seconds() > 2
}
