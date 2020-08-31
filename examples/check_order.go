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
	client.Referral = ""

	// Check an existing order
	order, err := client.CheckOrder(0000000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your order status:", order.Status)
	fmt.Println("SMS Code:", order.SMS[0].Code)

	/*OUTPUT
	  Your order status: FINISHED
	  SMS Code: 000000
	*/
}
