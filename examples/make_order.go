package main

import (
	"fmt"
	"log"

	"github.com/Fef0/go-5sim/fivesim"
)

/* WARNING THIS FILE SHOULD NOT BE EXECUTED AS IS*/
func main() {
	key := "yourKey"

	client := fivesim.NewClient(key)

	// Use this if you like the project and want to support my work
	client.Referral = "3a23a483"

	/* WARNING THIS FILE SHOULD NOT BE EXECUTED AS IS*/
	// Buy a new number from Russia with tele2 provider and available for Telegram activation
	// with no forwarding number
	order, err := client.BuyActivationNumber("russia", "tele2", "telegram", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your order status:", order.Status)

	/* OUTPUT
	   Your order status: PENDING
	*/
}
