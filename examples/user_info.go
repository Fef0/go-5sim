package main

import (
	"fmt"
	"log"

	"github.com/Fef0/go-5sim/fivesim"
)

func main() {
	key := "yourKey"

	client := fivesim.NewClient(key)

	// Use this if you like the project and want to support my work
	client.Referral = "3a23a483"

	// You can use this method to get a cumulative info about your user
	// using only one API call
	info, err := client.GetUserInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your ID is:", info.ID)
	fmt.Println("Your Email is:", info.Email)
	fmt.Println("Your Balance is:", info.Balance)
	fmt.Println("Your Rating is:", info.Rating)
	fmt.Println()

	// Or you can use single API calls for each info field
	userID, err := client.GetID()
	if err != nil {
		log.Fatal(err)
	}

	email, err := client.GetEmail()
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	}

	rating, err := client.GetRating()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your ID is:", userID)
	fmt.Println("Your Email is:", email)
	fmt.Println("Your Balance is:", balance)
	fmt.Println("Your Rating is:", rating)

	/*OUTPUT
	Your ID is: 00000
	Your Email is: email@example.in
	Your Balance is: 10.00
	Your Rating is: 95

	Your ID is: 00000
	Your Email is: email@example.in
	Your Balance is: 10.00
	Your Rating is: 95
	*/
}
