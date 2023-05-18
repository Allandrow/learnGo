package main

import "fmt"

func main()  {
	const totalTickets uint16 = 100
	
	var conferenceName string
	var remainingTickets uint16 = totalTickets
	var userName string
	var userEmail string
	var userTickets uint16
	var plannedUsers []string

	fmt.Println("Hello and welcome to the conference ticket booking application")
	fmt.Println("Which conference do you wish to buy tickets for?")
	fmt.Scan(&conferenceName)
	fmt.Printf(
		"%v has %v remaining tickets out of %v. How many do you wish to buy?\n",
		conferenceName,
		remainingTickets,
		totalTickets,
	)
	fmt.Scan(&userTickets)
	fmt.Printf("Great, we will put %v tickets aside for you. We will now need some basic information to reserve your tickets. What is your name?\n", userTickets)
	fmt.Scan(&userName)
	fmt.Printf("Nice to meet you %v, could you give us your email adress as well?\n", userName)
	fmt.Scan(&userEmail)
	fmt.Printf("All good %v! Your %v tickets are reserved for the %v conference under your name.\n", userName, userTickets, conferenceName)
	fmt.Printf("You will get a confirmation email sent to %v shortly.\n", userEmail)
	fmt.Println("Thanks you for using this app!")

	plannedUsers = append(plannedUsers, userName)

	remainingTickets -= userTickets
	fmt.Printf("The %v now has %v tickets still available.\n", conferenceName, remainingTickets)
	fmt.Printf("List of people coming to %v: %v\n", conferenceName, plannedUsers)
}