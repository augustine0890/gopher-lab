package main

import (
	"flag"
	"log"
	"time"
)

// Given a target date implement a function that outputs the number of nights until the given date.
var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	timeFormated, err := time.Parse(expectedFormat, target)
	if err != nil {
		log.Fatalf("Error to parsing the string")
	}

	if time.Now().After(timeFormated) {
		log.Fatalf("invalid date")
	}

	return timeFormated
}

// calcSleeps returns the number of sleeps until the target.
func calcSleep(target time.Time) float64 {
	timeSleep := time.Until(target).Hours() / 24
	return timeSleep
}

// go run sleeps.go -bday=YourBirthDay
func main() {
	bday := flag.String("bday", "", "Your next birthday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!", int(calcSleep(target)))
}
