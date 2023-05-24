package main

import (
	"fmt"
	"strings"
)

// Prompt for email and validate
// Add firstName, lastName and email in a struct, append to slice of attendees in city struct
// Display confirmation and start again.

func main() {
	GreetUser()

	var attendee = Attendee{}

	conference := getUserConferenceSelection()
	displayConferenceAttendees(conference)
	
	attendee.tickets = getUserTickets(conference.remainingTickets)
	conference.remainingTickets -= attendee.tickets
	fmt.Printf("%v tickets remaining out of %v\n", conference.remainingTickets, conference.totalTickets)

	fmt.Println("We will need some basic information to complete your reservation:")

	attendee.firstName = getUserFirstName()
	attendee.lastName = getUserLastName()
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

func getUserTickets(remainingTickets int16) int16 {
	fmt.Println("How many tickets do you wish to get?")
	var tickets int16

	for {
		fmt.Scan(&tickets)

		if isValidAmount(tickets, remainingTickets) {
			break
		}

		fmt.Println("Incorrect value, please retry.")
	}

	fmt.Printf("Great, we'll put %v tickets aside for you!\n", tickets)

	return tickets
}

func getUserFirstName() string {
	fmt.Println("What is your first name?")
	var name string

	for {
		fmt.Scan(&name)

		if isValidName(name) {
			break
		}

		fmt.Println("Incorrect value, please try again.")
	}

	return name
}

func getUserLastName() string {
	fmt.Println("What is your last name?")
	var name string

	for {
		fmt.Scan(&name)

		if isValidName(name) {
			break
		}

		fmt.Println("Incorrect value, please try again.")
	}

	return name
}