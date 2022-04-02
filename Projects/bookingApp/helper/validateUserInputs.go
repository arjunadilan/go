package helper // defines which package this belong to
import "strings"

func ValidateUserInputs(firstName, lastName, email string, userTickets, remainingTickets uint) (isValidName, isValidEmail, isValidTicket bool) {
	isValidName = len(firstName) > 2 && len(lastName) > 2
	isValidEmail = strings.Contains(email, "@")
	isValidTicket = userTickets <= remainingTickets
	return
}
