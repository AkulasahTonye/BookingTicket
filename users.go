package main

import "strings"

func ValidateUserInput(FirstName string, LastName string, userTickets uint, Email string) (bool, bool, bool)  {
	isValidName  := len(FirstName) >2 && len(LastName) >=2
	isValidEmail := strings.Contains(Email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
