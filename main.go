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

    fmt.Printf("Welcome to %v booking application\n", eventName)
    fmt.Printf("We have total of %v tickets and %v tickets are remaining.\n", eventTickets, remainingTickets)
    fmt.Printf("Get your tickets here to attend\n")

    for {

        var firstName string
        var lastName string
        var emailAddress string
        var userTickets uint
        //ask user for their name
        fmt.Println("Enter your first name: ")
        fmt.Scan(&firstName)
    
        fmt.Println("Enter your last name: ")
        fmt.Scan(&lastName)
    
        fmt.Println("Enter your email address: ")
        fmt.Scan(&emailAddress)
    
        fmt.Println("Enter number of tickets: ")
        fmt.Scan(&userTickets)
    
        remainingTickets = remainingTickets - userTickets
        bookings = append(bookings, firstName + " " + lastName)
    
        fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, emailAddress)
        fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
    
        firstNames := []string{}
        for _, booking := range bookings {
            var names = strings.Fields(booking)
            firstNames = append(firstNames, names[0])
        }
        fmt.Printf("The first names of bookings are: %v\n", firstNames)

    }


}

 