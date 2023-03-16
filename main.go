package main

import (
	"fmt"
	"strings"
)

func main(){
	confName := "Go Conference"
	const conftickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n",confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conftickets,remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")	

for{
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// ask user for their names
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of Tickets: ")
		fmt.Scan(&userTickets)
		if userTickets <= int(remainingTickets){
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName + " " + lastName )
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n",firstName,lastName,userTickets,email);
			fmt.Printf("%v tickets remaining for %v \n",remainingTickets,confName)
	
			firstNames := [] string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			} 
			fmt.Printf("These are all our bookings: %v \n",firstNames)
			if remainingTickets == 0{
				// End program
				fmt.Println("Our Conference is booked out come back next year.")
				break
			}
		}else{
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",remainingTickets,userTickets)
		}
	}
}