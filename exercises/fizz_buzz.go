package main

import "fmt"

func fizzBuzz(n int) {
	for num := 1; num <= n; num++ {
		switch {
		case num%5 == 0 && num%3 == 0:
			fmt.Println("FizzBuzz")
		case num%3 == 0:
			fmt.Println("Fizz")
		case num%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(num)
		}
	}
}

func main() {
	fizzBuzz(20)
}
