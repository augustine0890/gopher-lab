package main

import (
	"fmt"
	"strconv"
)

func isEvenEnded(num int) bool {
	strNum := strconv.Itoa(num)
	return strNum[0] == strNum[len(strNum)-1]
}

func countEvenEnded() int {
	count := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			if isEvenEnded(j * i) {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countEvenEnded())
}
