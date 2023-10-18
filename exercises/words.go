package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}
	return wordCount
}

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	wordCount := wordFrequency(text)
	fmt.Println(wordCount)
}
