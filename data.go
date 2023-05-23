package main

type Attendee struct {
	firstName string
	lastName string
	email string
	tickets int16
}

type Conference struct {
	name string
	totalTickets int16
	remainingTickets int16
	attendees []Attendee
}

var Conferences = []Conference{
	{
		name: "London",
		totalTickets: 100,
		remainingTickets: 100,
		attendees: make([]Attendee, 0),
	},
	{
		name: "Paris",
		totalTickets: 200,
		remainingTickets: 200,
		attendees: make([]Attendee, 0),
	},
	{
		name: "Tokyo",
		totalTickets: 80,
		remainingTickets: 80,
		attendees: make([]Attendee, 0),
	},
	{
		name: "Seoul",
		totalTickets: 150,
		remainingTickets: 150,
		attendees: make([]Attendee, 0),
	},
	{
		name: "Berlin",
		totalTickets: 50,
		remainingTickets: 50,
		attendees: make([]Attendee, 0),
	},
}