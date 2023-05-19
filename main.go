package main

import (
	"fmt"
	"strings"
)

func main()  {
	const totalTickets = 100
	
	var remainingTickets int16 = totalTickets
	var plannedUsers []string

	fmt.Println("Hello and welcome to the conference ticket booking application")

	conferenceName := getConferenceName()

	for remainingTickets > 0 {
		fmt.Printf(
			"%v has %v remaining tickets out of %v. How many do you wish to buy?\n",
			conferenceName,
			remainingTickets,
			totalTickets,
		)

		userTickets := handleTicketReservation(&remainingTickets)
		userName := getName()
		firstName := strings.Fields(userName)[0]
		userEmail := getValidatedEmail(firstName)		
		
		fmt.Printf("All good %v! Your %v tickets are reserved for the %v conference under your name.\n", firstName, userTickets, conferenceName)
		fmt.Printf("You will get a confirmation email sent to %v shortly.\n", userEmail)
		fmt.Println("Thanks you for using this app!")
	
		plannedUsers = append(plannedUsers, userName)
	
		fmt.Printf("The %v now has %v tickets still available.\n", conferenceName, remainingTickets)

		conferenceGoers := []string{}

		for _, plannedUser := range plannedUsers {
			var names = strings.Fields(plannedUser)
			initial := names[1][0:1]
			conferenceGoers = append(conferenceGoers, names[0] + " " + strings.ToUpper(initial) + ".")
		}

		fmt.Printf("List of people coming to %v: %v\n\n", conferenceName, conferenceGoers)
	}

	fmt.Printf("%v is fully booked, see you next time!\n", conferenceName)
}

func getConferenceName() string  {
	var conferenceName string
	
	for !isValidNameLength(conferenceName) {
		fmt.Println("Which conference do you wish to buy tickets for?")
		fmt.Scan(&conferenceName)

		if !isValidNameLength(conferenceName) {
			fmt.Println("Incorrect name, please re-enter")
		}

	}
	return conferenceName
}

func handleTicketReservation(remaining *int16) int16 {
	var userTickets int16
	for !isValidAmount(userTickets, *remaining) {		
		fmt.Scan(&userTickets)
		
		if !isValidAmount(userTickets, *remaining) {
			fmt.Println("Incorrect amount, please select another amount")
		}
	}

	fmt.Printf("Great, we will put %v tickets aside for you.\n", userTickets)
	*remaining -= userTickets

	return userTickets
}

func isValidAmount(value int16, limit int16) bool {
	return value > 0 && value <= limit
}

func getName() string {
	firstName := getValidatedName("first")
	lastName := getValidatedName("last")

	return firstName + " " + lastName
} 

func getValidatedName(key string) string {
	var name string
	for !isValidNameLength(name) {
		fmt.Printf("We will now need some basic information to reserve your tickets. What is your %v name?\n", key)
		fmt.Scan(&name)

		if !isValidNameLength(name) {
			fmt.Printf("Invalid name, please provide your %v name again\n", key)
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