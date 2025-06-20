package main

import "fmt"

func main() {
    eventName := "Earth Dance"
    const eventTickets int = 500
    var remainingTickets uint = 500

    fmt.Printf("Welcome to %v booking application\n", eventName)
    fmt.Printf("We have total of %v tickets and %v tickets are remaining.\n", eventTickets, remainingTickets)
    fmt.Printf("Get your tickets here to attend\n")


    var firstName string
    var lastName string
    var emailAddress string
    var userTickets int
    //ask user for their name
    fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email address: ")
    fmt.Scan(&emailAddress)

    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, emailAddress)
}

 