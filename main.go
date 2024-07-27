package main

import "fmt"

func main() {
	ConferenceName := "HNG Conference"
	const ConferenceTickets = 50
	var RemainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v Tickets and %v are still available.\nGet your tickets here to attend\n", ConferenceName, ConferenceTickets, RemainingTickets)

	var FirstName string
	var LastName string
	var Email string
	var userTickets uint

	//	Asking users for their name
	fmt.Println("Enter your First Name:")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&LastName)

	fmt.Println("Enter number of Tickets:")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your Email Address:")
	fmt.Scan(&Email)

	RemainingTickets = RemainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. you will recive a confirmation email at %v\n", FirstName, LastName, userTickets, Email)

	fmt.Printf("%v Tickets Remaining for %v\n", RemainingTickets, ConferenceName)
}
