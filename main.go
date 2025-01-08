package main

import ( "fmt"
	"strings" 
)


func main(){

	// another way of declaring variable is to use := syntax instead of delacring var and datatype
	// but this syntax won't work with constent variable and any explicit declared variable like unsigned variable below
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 
	// Arrays syntax
	// var bookings [50]string
	// Slice syntax
	var bookings []string

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avialable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket to attend the conference")
	
	for {
		// booking[0] = "Dawa"
		// booking[1] = "Nyima"

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// How to use input in golang and understand the pointer
		fmt.Println("Enter your First Name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last Name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("How many tickets you wanna book: ")
		fmt.Scan(&userTickets)

		// Check how the actual values and pointer stores in the memory
		// fmt.Println(conferenceName)
		// fmt.Println(&conferenceName)

		// fmt.Println(userName, userTickets)
		
		remainingTickets = remainingTickets - userTickets	

		// bookings[0] = firstName + "  " + lastName
		//Slice syntax prefix yeah use append module
		bookings = append(bookings, firstName + "  " + lastName)

		// fmt.Printf("The whole slice %v\n", bookings)
		// fmt.Printf("The datatype of the slice %T\n", bookings)
		// fmt.Printf("The value of the first slice %v\n", bookings[0])
		// fmt.Printf("The size of the slice %v\n", len(bookings))

		fmt.Printf("These are all our bookings %v\n", bookings)

		fmt.Printf("Thanks %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])

		}
		fmt.Printf("The first name of booking are: %v\n", firstNames)
	}

}
