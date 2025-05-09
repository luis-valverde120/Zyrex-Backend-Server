package models

/*
* Procinve Represents a state in the e-commerce application.
* It contains fields for the state ID and name.
 */
type Province struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
