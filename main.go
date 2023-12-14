package main

import (
	"fmt"
	"time"
)

func Greetings() string {
	timeNow := time.Now()
	return fmt.Sprintf("Hello, welcome to the conference the time is %s \n", timeNow.Format(time.RFC822))
}

func Onboarding(name *string, email *string, location *string, ticket *int) map[string]string {
	bioData := map[string]string{"name": *name, "email": *email, "location": *location, "ticket": string(rune(*ticket))}
	return bioData

}
func main() {
	var conferenceLocation string
	var ticketAvailable int = 30
	var ticketSold int
	var firstName string
	var email string

	greet := Greetings()
	fmt.Printf("%s", greet)

	println("what is your first name")
	fmt.Scan(&firstName)
	println("what is your email")
	fmt.Scan(&email)
	println("which conference do you wish to attend?")
	fmt.Scan(&conferenceLocation)
	println("How many tickets do you wish to buy")
	fmt.Scan(&ticketSold)
	fmt.Printf("Tickets available %d", (ticketAvailable - ticketSold))

	data := Onboarding(&firstName, &email, &conferenceLocation, &ticketSold)
	for key, value := range data {
		fmt.Printf("The key %s and value %s \n", key, value)

	}
	fmt.Printf("Your name is %s and your email is %s \n", firstName, email)

}
