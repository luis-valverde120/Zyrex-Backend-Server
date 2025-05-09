package models

import (
	"errors"
	"strings"
)

/*
* City Represents a city in the e-commerce application.
* It contains fields for the city ID and name.
 */
type City struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (city *City) Validate() error {
	if strings.TrimSpace(city.Name) == "" {
		return errors.New("city name is required")
	}
	return nil
}
