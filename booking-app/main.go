package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		GreetUser()

		var attendee = Attendee{}

		conference := getUserConferenceSelection()
	
		displayConferenceAttendees(conference)
		
		if conference.remainingTickets == 0 {
			fmt.Println("This conference is fully booked, sorry !")
			continue
		}
		
		attendee.tickets = getUserTickets(conference.remainingTickets)
		conference.remainingTickets -= attendee.tickets
	
		fmt.Println("We will need some basic information to complete your reservation:")
	
		attendee.firstName = getUserFirstName()
		attendee.lastName = getUserLastName()
		attendee.email = getUserEmail()
	
		conference.attendees = append(conference.attendees, attendee)
	
		fmt.Printf("All good %v! %v tickets have been reserved for the conference located in %v.\n", attendee.firstName, attendee.tickets, conference.city)
	}
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

func getUserConferenceSelection() *Conference {
	fmt.Println("Select the conference location you wish to attend:")
	var selection string

	for {
		fmt.Scan(&selection)

		for index, conference := range conferences {
			if strings.ToLower(conference.city) == strings.ToLower(selection) {
				return &conferences[index]
			}
		}

		printWrongValue()
	}
}

func displayConferenceAttendees(conference *Conference) {
	fmt.Printf("There are %v remaining tickets out of %v for the conference located in %v.\n", conference.remainingTickets, conference.totalTickets, conference.city)
	if len(conference.attendees) == 0 {
		fmt.Println("There are currently no attendees, be the first!")
	} else {
		for index, attendee := range conference.attendees {
			fmt.Printf("%v: %v with %v tickets\n", index + 1, attendee.firstName, attendee.tickets)
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

		printWrongValue()
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

		printWrongValue()
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

		printWrongValue()
	}

	return name
}

func getUserEmail() string {
	fmt.Println("What is your email adress?")
	var email string

	for {
		fmt.Scan(&email)

		if isValidEmail(email) {
			break
		}

		printWrongValue()
	}

	return email
}

func printWrongValue() {
	fmt.Println("Incorrect value, please try again.")
}