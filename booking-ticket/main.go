package main

import (
	"booking-app/helpers/array"
	"booking-app/helpers/input"
	"booking-app/types"
	"fmt"
	"sync"
	"time"
)

var goEvent = types.Event{
	Name:             "Go Conference",
	Tickets:          50,
	RemainingTickets: 50,
	Attendances:      make([]types.User, 0),
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for goEvent.RemainingTickets > 0 {
		//* Get user data and validate
		userName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketAmount := validateUserInput(userName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidTicketAmount {
			user := types.User{
				Name:            userName,
				Email:           userEmail,
				NumberOfTickets: userTickets,
			}

			//* Book and send tickets from user data
			bookTicket(user)
			wg.Add(1)
			go sendTicket(user)

			//* Print booking list
			names := getNames(goEvent.Attendances)
			array.FormattedPrint(names, "booking list: ")

			fmt.Printf("%d tickets remaining for %s.\n", goEvent.RemainingTickets, goEvent.Name)
		} else {
			var invalids []string
			if !isValidName {
				invalids = append(invalids, "name")
			}
			if !isValidEmail {
				invalids = append(invalids, "email")
			}
			if !isValidTicketAmount {
				invalids = append(invalids, "tickets amount")
			}

			//* Print invalid inputs
			array.FormattedPrint(invalids, "invalid inputs -> ")
			fmt.Println("Please booking again!")
		}
	}

	wg.Wait()
	fmt.Println("All tickets are sold out!")
}

func greetUsers() {
	fmt.Printf("Welcome to our %s booking application.\n", goEvent.Name)
	fmt.Printf("We have total of %d tickets and %d are still available.\n", goEvent.Tickets, goEvent.RemainingTickets)
}

func getUserInput() (name string, email string, amount uint) {
	fmt.Println("Enter your name and email:")
	if err := input.GetInput(&name, &email); err != nil {
		return
	}
	fmt.Println("Enter number of tickets:")
	if err := input.GetInput(&amount); err != nil {
		return
	}

	return
}

func validateUserInput(name string, email string, tickets uint) (bool, bool, bool) {
	return input.IsValidName(name), input.IsValidEmail(email), input.IsValidAmount(tickets, goEvent.RemainingTickets)
}

func bookTicket(user types.User) {
	goEvent.RemainingTickets -= user.NumberOfTickets
	goEvent.Attendances = append(goEvent.Attendances, user)
}

func getNames(users []types.User) []string {
	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}
	return names
}

func sendTicket(user types.User) {
	//* Set timer for sending
	time.Sleep(time.Second * 10)

	//* Print ticket
	ticket := fmt.Sprintf("%d tickets for %s", user.NumberOfTickets, user.Name)
	fmt.Println("--------------")
	fmt.Printf("Sending tickets:\n%s\nto email address %s\n", ticket, user.Email)
	fmt.Println("--------------")

	//* Decrease Wg counter when finished
	wg.Done()
}
