package models

type Customers struct {
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Dob          string `json:"dob"`
	Sex          int    `json:"sex"`
	Salt         string `json:"-"` // remove this items from json response
	Password     string `json:"-"`
	CreatedAt    string `json:"-"`
}
