package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"Cyril", "Nicolas", "Amaury", "Kévin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}