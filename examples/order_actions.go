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
	client.Referral = ""

	/* WARNING THIS FILE SHOULD NOT BE EXECUTED AS IS*/
	// Finish an order
	order, err := client.FinishOrder(0000000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your order status:", order.Status)

	/* WARNING THIS FILE SHOULD NOT BE EXECUTED AS IS*/
	// Cancel an order
	order, err = client.CancelOrder(0000000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your order status:", order.Status)

	/* WARNING THIS FILE SHOULD NOT BE EXECUTED AS IS*/
	// Ban an order
	order, err = client.BanOrder(0000000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your order status:", order.Status)

	/* OUTPUT
		   Your order status: FINISHED

	     Your order status: CANCELED

	     Your order status: BANNED
	*/
}
