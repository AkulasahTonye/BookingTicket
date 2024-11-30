package main


import (
	"fmt"
 	"FlightBookingTicket/users"
    "time"
)



const conferenceTickets = 50
var conferenceName = "HNG Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname string
	lastname string
	email string
	numberOfTickets uint
}



func main() {

	greetusers()

	for {
		firstname, lastname, email,   userTickets := GetUserInput()
	    isValidName, isValidEmail, isValidTicketNumber := users.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// Call out the function
			bookTicket(userTickets, firstname, lastname, email)
			go sendTickets(userTickets, firstname, lastname, email)

			firstNames := GetFirstName(  )
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			if  remainingTickets == 0{
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
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have total of %v Tickets and %v are still available.\n",conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")
}

func GetFirstName( ) []string  {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstname)
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
	remainingTickets = remainingTickets - userTickets

	var UserData = UserData{
		firstname: firstname,
		lastname: lastname,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, UserData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v Tickets Remaining for %v\n", remainingTickets, conferenceName);
}

func sendTickets(userTickets uint, firstname string, lastname string, email string)  {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v",userTickets, firstname, lastname)
	fmt.Println("###########")
	fmt.Printf("Sending Ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########")
}