package main

import "fmt"

func main()  {
	const totalTickets uint16 = 100
	
	var userName string
	var remainingTickets uint16 = totalTickets
	var email string
	var conferenceName string
	var wantedTickets uint16

	fmt.Println("Hello and welcome to the conference ticket booking application")
	fmt.Println("Which conference do you wish to buy tickets for?")
	fmt.Scan(&conferenceName)
	fmt.Printf(
		"%v has %v remaining tickets out of %v. How many do you wish to buy?\n",
		conferenceName,
		remainingTickets,
		totalTickets,
	)
	fmt.Scan(&wantedTickets)
	fmt.Printf("Great, we will put %v tickets aside for you. We will now need some basic information to reserve your tickets. What is your name?\n", wantedTickets)
	fmt.Scan(&userName)
	fmt.Printf("Nice to meet you %v, could you give us your email adress as well?\n", userName)
	fmt.Scan(&email)
	fmt.Printf("All good %v! Your %v tickets are reserved for the %v conference under your name.\n", userName, wantedTickets, conferenceName)
	fmt.Printf("You will get a confirmation email sent to %v shortly.\n", email)
	fmt.Println("Thanks you for using this app!")

	remainingTickets -= wantedTickets
	fmt.Printf("The %v now has %v tickets still available.\n", conferenceName, remainingTickets)
}