package main

import (
	"fmt"
	"strings"
)

type Guest struct {
	Name    string
	Email   string
	Tickets int
}

func main() {

	totalTicket := 50
	guests := []Guest{}

	for {

		if totalTicket == 0 {
			fmt.Println("🎉 The event is SOLD OUT! No more tickets available.")
			break
		}

		fmt.Println("\n===== Welcome to Olamide Event Ticket System =====")
		fmt.Println("Type 'exit' anytime to quit.")

		// NAME
		var name string
		fmt.Print("Enter name:  ")
		fmt.Scan(&name)

		if name == "exit" {
			fmt.Println("Exiting... Thanks for checking in!")
			break
		}

		if len(name) < 2 {
			fmt.Println("Name is too short. Try again.")
			continue
		}

		// AGE
		var age int
		fmt.Print("Enter your age: ")
		fmt.Scan(&age)

		if age < 18 {
			fmt.Println("Sorry, you're not eligible for this event.")
			continue
		}

		// EMAIL
		var email string
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email. Try again.")
			continue
		}

		// TICKETS
		var ticket int
		fmt.Printf("How many tickets do you want? Available: %d → ", totalTicket)
		fmt.Scan(&ticket)

		if ticket <= 0 {
			fmt.Println("You must buy at least 1 ticket.")
			continue
		}

		if ticket > totalTicket {
			fmt.Printf("Only %d tickets left. Try again.\n", totalTicket)
			continue
		}

		totalTicket -= ticket

		// SAVE GUEST
		guests = append(guests, Guest{Name: name, Email: email, Tickets: ticket})

		fmt.Println("\n🎫 Purchase Successful!")
		fmt.Println("----- Receipt -----")
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Tickets Bought:", ticket)
		fmt.Println("Tickets Remaining:", totalTicket)
		fmt.Println("-------------------")

		fmt.Println("\nCurrent Guests:", guests)
	}

}
