package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50
	var booking []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v Tickets and %v are still remaining\n", conferenceTicket, remainingTickets)
	fmt.Println("get your tickets here to attend your conference")

	for {

		var userName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("Enter your Name:")
		fmt.Scan(&userName)
		fmt.Println("Enter your LastName:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter Num of tickets:")
		fmt.Scan(&userTicket)
		if userTicket <= remainingTickets {
			

		remainingTickets = remainingTickets - userTicket
		// booking[0] = userName+" "+lastName
		booking = append(booking, userName+" "+lastName)

		fmt.Printf("The whole array %v:\n", booking)
		fmt.Printf("this is the first value : %v\n", booking[0])
		fmt.Printf("The type of array %T:\n", booking)
		fmt.Printf("The size of array %v:\n", len(booking))

		fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for  %v\n", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range booking {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are  all our first names of bookings: %v\n", firstNames)
		if remainingTickets == 0 {
			//End program
			fmt.Println("confrence tickets Soled Out")
			break
		}
		} else {
			fmt.Println("You have entered too many tickets, Try again ")
			
		}
	

}
}