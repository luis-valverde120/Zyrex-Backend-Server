package models

/*
* User Represents a user in the e-commerce application.
* It contains fields for the user ID, name, email, password, and address.
 */

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:gender`
}
