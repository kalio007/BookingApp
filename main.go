package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket = 50

	// "%v" is for referencing the value of variable that passed
	// "%T" is for referencing the type of the varible that goes into it.
	fmt.Println("Welcome to our", conferenceName, "Booking service")
	fmt.Println("we have total", conferenceTicket, "and remaining", remainingTicket)
	fmt.Println("Please get your ticket here for", conferenceTicket)
	fmt.Printf("Welcome to our %v \n", conferenceName)
	// collecting user's data

	var userName string
	fmt.Scan(userName)
	var ticket int
	fmt.Printf("user %v booked %v tickets. \n", userName, ticket)
}
