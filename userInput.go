package main

import "fmt"

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
	for !isValidName(name) {
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

func getUserName() string {
	fmt.Println("We will now need some basic information to reserve your tickets. What is your first name?")
	firstName := getValidatedName()
	
	fmt.Println("And your last name?")
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