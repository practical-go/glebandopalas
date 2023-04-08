package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello %s", name)
}

func twoPlusTwo() int {
	return 2 + 2
}

func main() {
	greet("Gleb!\n")

	fmt.Println("2 + 2: ", twoPlusTwo())
}
