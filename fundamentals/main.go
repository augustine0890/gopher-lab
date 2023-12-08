package main

import "fmt"

func main() {
	var a int
	a = 2
	var p *int
	p = &a
	fmt.Println(*p)
	a = 5
	fmt.Println(*p)
	*p = *p + 2
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	ages := map[string]int{
		"Peter": 42,
	}
	incrementPeterAge(ages)
	fmt.Println(ages)

	myMap := make(map[string]string)
	fmt.Println(myMap)
}

func incrementPeterAge(m map[string]int) {
	m["Peter"] += 1
}
