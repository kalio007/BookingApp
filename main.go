package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50

	// "%v" is for referencing the value of variable that passed
	// "%T" is for referencing the type of the varible that goes into it.
	fmt.Println("Welcome to our", conferenceName, "Booking service")
	fmt.Println("we have total", conferenceTicket, "and remaining", remainingTicket)
	fmt.Println("Please get your ticket here for", conferenceTicket)
	fmt.Printf("Welcome to our %v \n", conferenceName)

	//create an array for the bookings
	//array my have a size and its data type
	//this array is index based like
	// bookings[0] = "Nana"
	var bookings [50]string

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

	remainingTicket = remainingTicket - userTicket
	//to perfome ops the vairble has to have the same type

	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v \n", bookings)
	//Printing thr first person that booked for a price
	fmt.Printf("the first person: %v \n", bookings[0])
	//printing the type of the array
	fmt.Printf("the type of the array: %T \n", bookings)
	//printing the lenght of the array
	fmt.Printf("Array lenght: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v shortly.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v. \n", remainingTicket, conferenceName)

}
