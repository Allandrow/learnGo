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

func getUserInput() (string, int16, string, string) {
	conferenceName := getValidatedConferenceName()
	tickets := getConferenceTickets(conferenceName)
	userName := getUserName()
	userEmail := getUserEmail(userName)

	return conferenceName, tickets, userName, userEmail
}

func getValidatedConferenceName() string {
	var name string
	fmt.Println("Which conference do you wish to buy tickets for?")
	for isValidName(name) {
		fmt.Scan(&name)

		if !isValidName(name) {
			fmt.Println("Invalid conference name, retry.")
		}
	}

	return name
}

func getConferenceTickets(conference string) int16 {
	var tickets int16

	fmt.Printf(
		"%v has %v remaining tickets out of %v. How many do you wish to buy?\n",
		conference,
		remainingTickets,
		totalTickets,
	)
	
	for !isValidAmount(tickets, remainingTickets) {		
		fmt.Scan(&tickets)
		
		if !isValidAmount(tickets, remainingTickets) {
			fmt.Println("Incorrect amount, please select another amount")
		}
	}

	fmt.Printf("Great, we will put %v tickets aside for you.\n", tickets)
	remainingTickets -= tickets

	return tickets
}

func isValidAmount(value int16, limit int16) bool {
	return value > 0 && value <= limit
}

func getUserName() string {
	fmt.Println("We will now need some basic information to reserve your tickets. What is your first name?")
	firstName := getValidatedName()
	lastName := getValidatedName()

	return firstName + " " + lastName
}

func getValidatedName() string {
	var name string

	for !isValidName(name) {
		fmt.Scan(&name)

		if !isValidName(name) {
			fmt.Println("Incorrect value, please enter again")
		}
	}

	return name
}

func getUserEmail(name string) string {
	var email string
	
	for !isValidEmail(email){
		fmt.Printf("Nice to meet you %v, could you give us your email adress as well?\n", name)
		fmt.Scan(&email)

		if !isValidEmail(email) {
			fmt.Println("Incorrect email, please re-enter your email.")
		}
	}

	return email
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func isValidName(name string) bool {
	return len(name) > 2
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
