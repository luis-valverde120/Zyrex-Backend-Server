package models

type Store struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AddressID   int    `json:"address_id"`
	OwnerID     int    `json:"owner_id"`
}
