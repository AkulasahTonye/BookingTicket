package main

import "fmt"

func main() {
	var ConferenceName = "HNG"
	const ConferenceTickets = 50
	var RemainingTickets uint = 50

	RemainingTickets = 7

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v Tickets and %v are still available.\nGet your tickets here to attend\n", ConferenceName, ConferenceTickets, RemainingTickets)

	var userName string
	var userTicket int
	//	Asking users for their name
	userName = "Tom"
	userTicket = 2
	fmt.Printf("user %v booked %v tickits\n", userName, userTicket)

}
