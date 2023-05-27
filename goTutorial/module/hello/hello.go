package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("Cyril")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}