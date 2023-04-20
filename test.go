package main

import (
	"fmt"
)

var name string

func sayHi() {
	var name string
	fmt.Println("Hello what is your name?")
	fmt.Scan(&name)
	fmt.Printf("Hello %v", name)
}

func main() {
	sayHi()
}
