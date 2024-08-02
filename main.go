package main

import (
	"fmt"
	"strings"
)

func main() {
	ConferenceName := "HNG Conference"
	const ConferenceTickets = 50
	var RemainingTickets uint = 50
	var bookings []string
	
	greetusers(ConferenceName,ConferenceTickets,RemainingTickets)
	
	for {
		var FirstName string
		var LastName string
		var Email string
		var userTickets uint

		//	Asking users for their name
		fmt.Println("Enter your First Name:")
		fmt.Scan(&FirstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&LastName)

		fmt.Println("Enter your Email Address:")
		fmt.Scan(&Email)

		fmt.Println("Enter number of Tickets:")
		fmt.Scan(&userTickets)

	    isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(FirstName, LastName, userTickets, RemainingTickets, Email)

		if isValidName && isValidEmail && isValidTicketNumber {
			RemainingTickets = RemainingTickets - userTickets
			bookings = append(bookings, FirstName+" "+LastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", FirstName, LastName, userTickets, Email)

			fmt.Printf("%v Tickets Remaining for %v\n", RemainingTickets, ConferenceName)

			// Call out the function
			
			firstNames := GetFirstName(bookings) 
			fmt.Printf("The first name of bookings are: %v\n", firstNames)


			if  RemainingTickets == 0{
				//End program
				fmt.Println("Our Confrence is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is to short")
			}
			if !isValidEmail{
				fmt.Println("Your email Address doesn't contain the @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("The number of tickets you entered is invalid")
			}
		}
	}
	
}

func greetusers(confName string, confTicket int, RemainingTickets uint){
	fmt.Printf("Welcome to %v booking application. \n", confName)
	fmt.Printf("We have total of %v Tickets and %v are still available.\n",confTicket, RemainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")
}

func GetFirstName(bookings []string) []string  {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func ValidateUserInput(FirstName string, LastName string, userTickets uint, RemainingTickets uint, Email string) (bool, bool, bool)  {
	isValidName := len(FirstName) >2 && len(LastName) >= 2
	isValidEmail := strings.Contains(Email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
	
}
