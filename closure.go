package main

import "fmt"

func main() {
	name := "Chasa"
	counter := 0

	increment := func() {
		name := "Farel"
		fmt.Println("Increment")
		counter++

		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
