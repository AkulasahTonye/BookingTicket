package main

import (
	"fmt"
	"strings"
)

const ConferenceTickets = 50
var ConferenceName = "HNG Conference"
var RemainingTickets uint = 50
var bookings = []string{}



func main() {

	greetusers()

	for {
		firstname, lastname, email,   userTickets := GetUserInput()
	    isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstname, lastname, userTickets,  email)

		if isValidName && isValidEmail && isValidTicketNumber {

			// Call out the function
			bookTicket(userTickets, firstname, lastname, email)

			firstNames := GetFirstName(  )
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

func greetusers(){
	fmt.Printf("Welcome to %v booking application. \n", ConferenceName)
	fmt.Printf("We have total of %v Tickets and %v are still available.\n",ConferenceTickets, RemainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")
}

func GetFirstName( ) []string  {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func GetUserInput() (string, string, string, uint){
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	//	Asking users for their name
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstname)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your Email Address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of Tickets:")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTicket( userTickets uint, firstname string, lastname string, email string,)  {
	RemainingTickets = RemainingTickets - userTickets
	bookings = append(bookings, firstname+" "+lastname)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v Tickets Remaining for %v\n", RemainingTickets, ConferenceName);;
}



