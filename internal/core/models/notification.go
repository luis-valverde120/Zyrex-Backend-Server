package models

import (
	"errors"
	"strings"
	"time"
)

// NotificationType is a list of all the types of notifications.
type NotificationType string

const (
	NotificationOrder    NotificationType = "order"
	NotificationErrSync  NotificationType = "error_sync"
	NotificationLowStock NotificationType = "low_stock"
	NotificationInfo     NotificationType = "info"
)

/*
* Notify Represents a notification in the e-commerce application.
* It contains fields for the notification ID, owner ID, type, message, read status, and creation date.
 */
type Notification struct {
	ID         int              `json:"id" gorm:"primaryKey"`
	OwnerID    int              `json:"owner_id"`
	Type       NotificationType `json:"type"`
	Message    string           `json:"message"`
	ReadStatus bool             `json:"read_status"`
	CreatedAt  time.Time        `json:"created_at"`
}

// * methods for Notify struct

func (notification *Notification) Validate() error {
	if notification.OwnerID == 0 {
		return errors.New("owner ID is required")
	}
	if strings.TrimSpace(notification.Message) == "" {
		return errors.New("message is required")
	}

	return nil
}

// * IsValidType is a validator of the type field.
func (notification *Notification) IsValidNotificationType(notification_type string) bool {
	switch strings.ToLower(notification_type) {
	case string(NotificationOrder), string(NotificationErrSync), string(NotificationLowStock), string(NotificationInfo):
		return true
	}
	return false
}

// * MarkAsRead marks the notification as read.
func (notification *Notification) MarkAsRead() {
	notification.ReadStatus = true
}

// * IsRead checks if the notification is read.
func (notification *Notification) IsRead() bool {
	return notification.ReadStatus
}

// * method construnctor
func NewNotification(ownerID int, t NotificationType, msg string) (*Notification, error) {
	notification := &Notification{
		OwnerID:    ownerID,
		Type:       t,
		Message:    msg,
		ReadStatus: false,
		CreatedAt:  time.Now(),
	}
	if err := notification.Validate(); err != nil {
		return nil, err
	}
	return notification, nil
}
