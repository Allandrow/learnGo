package main

type Attendee struct {
	firstName string
	lastName string
	email string
	tickets int16
}

type Conference struct {
	city string
	totalTickets int16
	remainingTickets int16
	attendees []Attendee
}

var Conferences = []Conference{
	{
		city: "Berlin",
		totalTickets: 50,
		remainingTickets: 50,
		attendees: make([]Attendee, 0),
	},
	{
		city: "London",
		totalTickets: 100,
		remainingTickets: 100,
		attendees: make([]Attendee, 0),
	},
	{
		city: "Paris",
		totalTickets: 200,
		remainingTickets: 200,
		attendees: make([]Attendee, 0),
	},
	{
		city: "Seoul",
		totalTickets: 150,
		remainingTickets: 150,
		attendees: make([]Attendee, 0),
	},
	{
		city: "Tokyo",
		totalTickets: 80,
		remainingTickets: 80,
		attendees: make([]Attendee, 0),
	},	
}