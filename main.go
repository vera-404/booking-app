package main

import (
	"fmt"
   // "strconv"
    "time"
    "sync"
)

const eventTickets int = 500

var eventName = "Earth Dance"
var remainingTickets uint = 500
//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
    firstName string
    lastName string
    emailAddress string
    numberOfTickets uint
}

 var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {

    firstName, lastName, emailAddress, userTickets := getUserInput()
    isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

    if isValidName && isValidEmail && isValidTicketNumber {

        bookTicket(userTickets, firstName, lastName, emailAddress)

        wg.Add(1)
        go sendTicket(userTickets, firstName, lastName, emailAddress)

        firsties := getFirstNames()
        fmt.Printf("The first names of bookings are: %v\n", firsties)

        if remainingTickets == 0 {
            //end program
            fmt.Println("Our Event is booked out. Come back next year.")
            //break
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
	//}
    wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", eventName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining.\n", eventTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, book := range bookings {
		//firstNames = append(firstNames, book["firstName"])
        firstNames = append(firstNames, book.firstName)
	}
	return firstNames

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

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {

	remainingTickets = remainingTickets - userTickets

    //create a map for a user
    //var userData = make(map[string]string)
    var userData = UserData {
        firstName : firstName,
        lastName : lastName,
        emailAddress : emailAddress,
        numberOfTickets: userTickets,
    }
   // userData["firstName"] = firstName
    //userData["lastName"] = lastName
    //userData["email"] = emailAddress
    //userData["no of tickets"] = strconv.FormatUint(uint64(userTickets), 10)


	bookings = append(bookings, userData)
    fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
    time.Sleep(50 * time.Second)
    var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
    fmt.Println("###################")
    fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, emailAddress)
    fmt.Println("###################")
    wg.Done()
}