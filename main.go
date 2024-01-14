package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50

	// "%v" is for referencing the value of variable that passed
	// "%T" is for referencing the type of the varible that goes into it.
	fmt.Println("Welcome to our", conferenceName, "Booking service")
	fmt.Println("we have total", conferenceTicket, "and remaining", remainingTickets)
	fmt.Println("Please get your ticket here for", conferenceTicket)
	fmt.Printf("Welcome to our %v \n", conferenceName)

	// for for loop
	var remainingTicket uint = 50
	//slice
	var bookings []string

	for {
		//create an array for the bookings
		//array my have a size and its data type
		//this array is index based like
		// bookings[0] = "Nana"

		//var bookings [50]string

		// to convert the array to slice a more convienet array

		// collecting user's data
		var firstName string
		var lastName string
		var userName string
		var email string
		var userTicket uint
		fmt.Println("please enter your First name")
		fmt.Scan(&firstName)
		fmt.Println("please enter your Last name")
		fmt.Scan(&lastName)
		fmt.Println("please enter your email adress")
		fmt.Scan(&email)
		fmt.Println("please enter your preffered Username")
		fmt.Scan(&userName)
		fmt.Println("how many tickets would you be buying?")
		fmt.Scan(&userTicket)

		if userTicket <= remainingTicket {
			//using the array
			// bookings[0] = firstName + " " + lastName   -we use this for an array
			//to perfome ops the vairble has to have the same type
			remainingTicket = remainingTicket - userTicket
			//using the slice
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole Slice: %v \n", bookings)
			//Printing thr first person that booked for a price
			fmt.Printf("the first person: %v \n", bookings[0])
			//printing the type of the array
			fmt.Printf("the type of the Slice: %T \n", bookings)
			//printing the lenght of the array
			fmt.Printf("Slice lenght: %v \n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v shortly.\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tickets remaining for %v. \n", remainingTicket, conferenceName)

			firstNames := []string{}
			//to iterate through a list
			for _, booking := range bookings {
				//to remove the error of a variable we are not interested in uisng we use the "_"
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("the first names of bookings are: %v\n", firstNames)
			fmt.Printf("This ia all out current bookings %v.\n", bookings)

			var noTickets bool = remainingTicket == 0 //no need to save it in a variablethough but i am doing so to test bool
			if noTickets {
				fmt.Println("Our confernece tickets is sold out")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remianig, so you can't book %v tickets.\n", remainingTicket, userTicket)
		}
	}
}
