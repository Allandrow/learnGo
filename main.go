package main

import (
	"fmt"
	"strings"
)

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
	displayConferenceAttendees(conference)
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

func getUserConferenceSelection() Conference {
	fmt.Println("Select the conference location you wish to attend:")
	var selection string

	for {
		fmt.Scan(&selection)

		for _, conference := range conferences {
			if strings.ToLower(conference.city) == strings.ToLower(selection) {
				return conference
			}
		}

		fmt.Println("Incorrect location, please try again")
	}
}

func displayConferenceAttendees(conference Conference) {
	fmt.Printf("There are %v remaining tickets out of %v for the conference located in %v.\n", conference.remainingTickets, conference.totalTickets, conference.city)
	if len(conference.attendees) == 0 {
		fmt.Println("There are currently no attendees, be the first!")
	} else {
		for index, attendee := range conference.attendees {
			fmt.Printf("%v: %v with %v tickets", index + 1, attendee.firstName, attendee.tickets)
		}
	}
}