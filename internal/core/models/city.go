package models

/*
* City Represents a city in the e-commerce application.
* It contains fields for the city ID and name.
 */
type City struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
