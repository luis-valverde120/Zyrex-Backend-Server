package models

/*
* Owner Represents the owner of a product in the e-commerce application.
* It contains fields for the owner's ID, first name, last name, phone number, email, password, and
* Gender.
* This struct is used to store information about owners in the database.
 */
type Owner struct {
	User
	StoreID int `json:"store_id" gorm:"index"`
}
