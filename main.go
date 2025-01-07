package main

import "fmt"


func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50 

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avialable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you ticket to attend the conference")

	
}
