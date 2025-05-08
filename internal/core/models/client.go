package models

/*
* Client Represents a client in the e-commerce application.
* It contains fields for the client ID, first name, last name, phone number, email
* address, password
* This struct is used to store information about clients in the database.
 */
type Client struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	AddressID int    `json:"address_id"`
}
