package main

import "fmt"

// Prompt for city
// Validate selection, if no city found reiterate prompt
// If city, display list of attendees and remaining tickets
// Prompt for desired amount of tickets
// Validate amount, if above remaining or 0 or below, reiterate prompt
// substract amount from remaining for new remaining
// Prompt for firstName and validate
// Prompt for lastName and validate
// Prompt for email and validate
// Add firstName, lastName and email in a struct, append to slice of attendees in city struct
// Display confirmation and start again.

func main() {
	GreetUser()
}

func GreetUser() {
	fmt.Println("Hello and welcome. Here is the list of available conferences:")
	var index = 1
	for _, conference := range Conferences {
		if conference.remainingTickets > 0 {
			fmt.Printf("%v: %v\n", index, conference.city)
			index++
		}
	}
}