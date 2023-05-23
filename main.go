package main

import (
	"fmt"
	"strings"
)

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

	conference := getUserConferenceSelection()
	fmt.Print("Selection is: " + conference)
}

func GreetUser() {
	fmt.Println("Hello and welcome. Here is the list of available conferences:")
	var index = 1
	for _, conference := range conferences {
		if conference.remainingTickets > 0 {
			fmt.Printf("%v: %v\n", index, conference.city)
			index++
		}
	}
}

func getUserConferenceSelection() string {
	fmt.Println("Select the conference location you wish to attend:")
	var selection string

	for {
		fmt.Scan(&selection)

		for _, conference := range conferences {
			if strings.ToLower(conference.city) == strings.ToLower(selection) {
				return conference.city
			}
		}

		fmt.Println("Incorrect location, please try again")
	}

}