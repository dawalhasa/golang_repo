package main

import "fmt"


func main(){
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50 

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avialable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket to attend the conference")

	var userName string
	var userTickets int



	userName = "Tom"
	userTickets = 2

	fmt.Println(userName, userTickets)
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	
}
