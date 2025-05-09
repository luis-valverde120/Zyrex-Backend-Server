package models

import "errors"

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

// * Methods for Address struct

// GetID returns the ID of the address.
func (address *Address) GetID() int {
	return address.ID
}

// GetStreet1 returns the first line of the street address.
func (address *Address) ChangeStreet1(street1 string) {
	address.Street_1 = street1
}

// ChangeStreet2 changes the second street address of the address.
func (address *Address) ChangeStreet2(street2 string) {
	address.Street_2 = street2
}

// ChangeCityID changes the city ID of the address.
func (address *Address) ChangeCityID(cityID string) {
	address.CityID = cityID
}

// ChangeStateID changes the state ID of the address.
func (address *Address) ChangeStateID(stateID string) {
	address.StateID = stateID
}

func (address *Address) Validate() error {
	if address.Street_1 == "" {
		return errors.New("street_1 is required")
	}
	// address 2 is optional, so no validation needed
	if address.CityID == "" {
		return errors.New("city_id is required")
	}
	if address.StateID == "" {
		return errors.New("state_id is required")
	}
	return nil
}
