package models

/*
* Notify Represents a notification in the e-commerce application.
* It contains fields for the notification ID, owner ID, type, message, read status, and creation date.
 */
type Notify struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	OwnerID    int    `json:"owner_id"`
	Type       int    `json:"type"`
	Message    string `json:"message"`
	ReadStatus bool   `json:"read_status"`
	CreatedAt  string `json:"created_at"`
}
