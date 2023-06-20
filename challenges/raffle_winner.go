package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

// Given an input.json file with raffle entries implement a function that outputs a list of
// structs containing the entries.

const path = "entries.json"

// raffleEntry is the struct we unmarshall raffle entries into
type raffleEntry struct {
	Id      string
	Name    string
	Country string
}

// importData reads the raffle entries from file and creates the entries slice
func importData() []raffleEntry {
	// Open the file
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error opening (reading) file: %v", err)
	}

	// // Read the file's content into a byte slice
	// data, err := io.ReadAll(file)
	// if err != nil {
	// log.Fatalf("Error reading file: %v", err)
	// }

	var entries []raffleEntry
	err = json.Unmarshal(file, &entries)
	if err != nil {
		log.Fatalf("Error unmarshalling: %v", err)
	}
	return entries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
