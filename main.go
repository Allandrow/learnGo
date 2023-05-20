package main

import (
	"fmt"
	"strings"
)

func main()  {
	const totalTickets = 100
	var remainingTickets int16 = totalTickets
	var plannedUsers []string
	var userTickets int16 = 0

	fmt.Println("Hello and welcome to the conference ticket booking application")

	conferenceName := getValidatedName("conference")

	for remainingTickets > 0 {
		fmt.Printf(
			"%v has %v remaining tickets out of %v. How many do you wish to buy?\n",
			conferenceName,
			remainingTickets,
			totalTickets,
		)

		userTickets = handleTicketReservation(remainingTickets)
		remainingTickets -= userTickets

		firstName := getValidatedName("first")
		lastName := getValidatedName("last")
		userEmail := getValidatedEmail(firstName)		
		
		endValidationMessage(firstName, userTickets, conferenceName, userEmail)
	
		plannedUsers = append(plannedUsers, firstName + " " + lastName)
		listAttendees(conferenceName, plannedUsers)
	}

	fmt.Printf("%v is fully booked, see you next time!\n", conferenceName)
}

func handleTicketReservation(remaining int16) int16 {
	var userTickets int16
	for !isValidAmount(userTickets, remaining) {		
		fmt.Scan(&userTickets)
		
		if !isValidAmount(userTickets, remaining) {
			fmt.Println("Incorrect amount, please select another amount")
		}
	}

	fmt.Printf("Great, we will put %v tickets aside for you.\n", userTickets)

	return userTickets
}

func isValidAmount(value int16, limit int16) bool {
	return value > 0 && value <= limit
}

func getValidatedName(key string) string {
	switch key {
		case "conference":
			fmt.Println("Which conference do you wish to buy tickets for?")
		case "first":
			fmt.Println("We will now need some basic information to reserve your tickets. What is your first name?")
		case "last":
			fmt.Println("And your last name?")
		default:
			fmt.Println("Something wrong happened")
		}

	return validateName()
}

func validateName() string {
	var name string

	for !isValidNameLength(name) {
		fmt.Scan(&name)

		if !isValidNameLength(name) {
			fmt.Println("Incorrect value, please enter again")
		}
	}

	return name
}

func isValidNameLength(name string) bool {
	return len(name) > 2
}

func getValidatedEmail(name string) string {
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

func concatWithInitials(name []string) string {
	initial := name[1][0:1]
	return name[0] + " " + strings.ToUpper(initial) + "."
}

func setConferenceAttendees(plannedUsers []string) []string {
	attendees := []string{}
		for _, plannedUser := range plannedUsers {
			fullName := strings.Fields(plannedUser)
			attendees = append(attendees, concatWithInitials(fullName))
		}

		return attendees
}

func listAttendees(conference string, attendees []string) {
	conferenceAttendees := setConferenceAttendees(attendees)

	fmt.Printf("List of people coming to %v:\n", conference)
	for index, attendee := range conferenceAttendees {
		fmt.Printf("%v: %v\n", index+1, attendee)
	}
}

func endValidationMessage(name string, tickets int16, conference string, email string) {
	fmt.Printf("All good %v! Your %v tickets are reserved for the %v conference under your name.\n", name, tickets, conference)
		fmt.Printf("You will get a confirmation email sent to %v shortly.\n", email)
		fmt.Println("Thanks you for using this app!")
}