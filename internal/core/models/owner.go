package models

/*
* Owner Represents the owner of a product in the e-commerce application.
* It contains fields for the owner's ID, first name, last name, phone number, email, password, and
* Gender.
* This struct is used to store information about owners in the database.
 */
type Owner struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
}
