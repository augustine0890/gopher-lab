package main

import (
	"log"
	"strings"
	"time"
)

// Implement a function that takes in a string and outputs a prolonged verion of each word

const deplay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(deplay)
}

// slowDown takes the given string and repeats its characters according to their index in the string
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	for _, word := range words {
		var str []string
		for i, ch := range word {
			str = append(str, strings.Repeat(string(ch), i+1))
		}

		print(strings.Join(str, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
