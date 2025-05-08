package models

/*
* Address Represents an address in the system.
* It contains fields for street address, city, state, and other relevant information.
 */
type Address struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Street_1 string `json:"street_1"`
	Street_2 string `json:"street_2"`
	CityID   string `json:"city_id"`
	StateID  string `json:"state_id"`
}
