package main

import (
	"fmt"
	"strings"
)

const totalTickets = 100
var remainingTickets int16 = totalTickets
var attendees []string

func main()  {
	for remainingTickets > 0 {
		greetUser()
	
		conference, tickets, userName, userEmail := getUserInput()
	
		printValidationMessage(conference, tickets, userName, userEmail)
		attendees = append(attendees, userName)
		printAttendeeList(conference)

		if remainingTickets == 0 {
			fmt.Printf("%v is fully booked, see you next time!\n", conference)
			break
		}
	}
}

func greetUser() {
	fmt.Println("Hello and welcome to the conference ticket booking application")
}

func printValidationMessage(conference string, tickets int16, userName string, email string) {
	fmt.Printf("All good %v! Your %v tickets are reserved for the %v conference under your name.\n", userName, tickets, conference)
	fmt.Printf("You will get a confirmation email sent to %v shortly.\n", email)
	fmt.Println("Thanks you for using this app!")
}

func printAttendeeList(conference string) {
	conferenceAttendees := setConferenceAttendees(attendees)

	fmt.Printf("List of people coming to %v:\n", conference)
	for index, attendee := range conferenceAttendees {
		fmt.Printf("%v: %v\n", index + 1, attendee)
	}
}

func setConferenceAttendees(plannedUsers []string) []string {
	attendees := []string{}
		for _, plannedUser := range plannedUsers {
			fullName := strings.Fields(plannedUser)
			attendees = append(attendees, concatWithInitials(fullName))
		}

		return attendees
}

func concatWithInitials(name []string) string {
	initial := name[1][0:1]
	return name[0] + " " + strings.ToUpper(initial) + "."
}
