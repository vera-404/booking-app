package main

import (
	"fmt"
	"strings"
)

func main() {
	eventName := "Earth Dance"
	const eventTickets int = 500
	var remainingTickets uint = 500
	bookings := []string{}

	greetUsers(eventName, eventTickets, remainingTickets)

	for {

		firstName, lastName, emailAddress, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, emailAddress, eventName)

			firsties := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firsties)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Event is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address does not conatin @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining.\n", confTickets, remTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames(books []string) []string {
	firstNames := []string{}
	for _, book := range books {
		var names = strings.Fields(book)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets

}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, emailAddress string, eventName string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}
