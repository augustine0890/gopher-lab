package main

import (
	"encoding/json"
	"log"
	"os"
)

// Given a list, users and their details implement a function that outputs the country with the most users

// User represents a user record
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	return data
}

// getBiggestMarket takes in the slice of users and returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	var markets = make(map[string]int)

	for _, user := range users {
		markets[user.Country] += 1
	}
	var market string
	var maxValue int
	for m, v := range markets {
		if v > maxValue {
			market = m
			maxValue = v
		}
	}
	return market, maxValue
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n", country, count)
}
