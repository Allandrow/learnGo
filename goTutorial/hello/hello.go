package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Cyril")
	fmt.Println(message)
}