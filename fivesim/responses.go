package fivesim

import "time"

// Product represents a single anonymous product
type Product struct {
	Category string  `json:"Category"`
	Quantity int     `json:"Qty"`
	Price    float32 `json:"Price"`
}

// Products represents a map of servicename:product
type Products map[string]Product

// UserInfo represents info about the user
type UserInfo struct {
	ID      int     `json:"id"`
	Email   string  `json:"email"`
	Balance float32 `json:"balance"`
	Rating  float32  `json:"rating"`
}

// SMS represents info about an incoming SMS
type SMS struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Date      time.Time `json:"date"`
	Sender    string    `json:"sender"`
	Text      string    `json:"text"`
	Code      string    `json:"code"`
}

// ActivationOrder represents info about an activation order
type ActivationOrder struct {
	ID               int       `json:"id"`
	Phone            string    `json:"phone"`
	Operator         string    `json:"operator"`
	Product          string    `json:"product"`
	Price            float32   `json:"price"`
	Status           string    `json:"status"`
	Expires          time.Time `json:"expires"`
	SMS              []SMS     `json:"sms"`
	CreatedAt        time.Time `json:"created_at"`
	Forwarding       bool      `json:"forwarding"`
	ForwardingNumber string    `json:"forwarding_number"`
	Country          string    `json:"russia"`
}

// HostingOrder represents info about an hosting order
type HostingOrder struct {
	ID      int       `json:"id"`
	Phone   string    `json:"phone"`
	Product string    `json:"product"`
	Price   float32   `json:"price"`
	Status  string    `json:"status"`
	Expires time.Time `json:"expires"`
	SMS     []SMS     `json:"sms"`
}
