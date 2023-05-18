package main

import (
	"fmt"
	"strings"
)

func main()  {
	const totalTickets uint16 = 100
	
	var conferenceName string
	var remainingTickets uint16 = totalTickets
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint16
	var plannedUsers []string

	fmt.Println("Hello and welcome to the conference ticket booking application")

	for {
		fmt.Println("Which conference do you wish to buy tickets for?")
		fmt.Scan(&conferenceName)

		for {
			fmt.Printf(
				"%v has %v remaining tickets out of %v. How many do you wish to buy?\n",
				conferenceName,
				remainingTickets,
				totalTickets,
			)
			
			fmt.Scan(&userTickets)
			
			if userTickets <= remainingTickets {
				break
			}
			
			fmt.Println("Not enough tickets available, please try to get less tickets")
		}

		fmt.Printf("Great, we will put %v tickets aside for you.\n", userTickets)
		remainingTickets -= userTickets
		
		fmt.Println("We will now need some basic information to reserve your tickets. What is your first name?")
		fmt.Scan(&firstName)
		
		fmt.Println("And your last name ?")
		fmt.Scan(&lastName)
		
		fmt.Printf("Nice to meet you %v %v, could you give us your email adress as well?\n", firstName, lastName)
		fmt.Scan(&userEmail)
		
		fmt.Printf("All good %v! Your %v tickets are reserved for the %v conference under your name.\n", firstName, userTickets, conferenceName)
		fmt.Printf("You will get a confirmation email sent to %v shortly.\n", userEmail)
		fmt.Println("Thanks you for using this app!")
	
		plannedUsers = append(plannedUsers, firstName + " " + lastName)
	
		fmt.Printf("The %v now has %v tickets still available.\n", conferenceName, remainingTickets)

		conferenceGoers := []string{}

		for _, plannedUser := range plannedUsers {
			var names = strings.Fields(plannedUser)
			initial := names[1][0:1]
			conferenceGoers = append(conferenceGoers, names[0] + " " + strings.ToUpper(initial) + ".")
		}

		fmt.Printf("List of people coming to %v: %v\n\n", conferenceName, conferenceGoers)

		if remainingTickets == 0 {
			fmt.Printf("%v is fully booked, see you next time!\n", conferenceName)
			break
		}
	}
}