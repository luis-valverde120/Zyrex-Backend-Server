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
type Notify struct {
	ID         int              `json:"id" gorm:"primaryKey"`
	OwnerID    int              `json:"owner_id"`
	Type       NotificationType `json:"type"`
	Message    string           `json:"message"`
	ReadStatus bool             `json:"read_status"`
	CreatedAt  time.Time        `json:"created_at"`
}

// * methods for Notify struct

func (notify *Notify) Validate() error {
	if notify.OwnerID == 0 {
		return errors.New("owner ID is required")
	}
	if strings.TrimSpace(notify.Message) == "" {
		return errors.New("message is required")
	}

	return nil
}

// * IsValidType is a validator of the type field.
func (notify *Notify) IsValidNotificationType(notification_type string) bool {
	switch strings.ToLower(notification_type) {
	case string(NotificationOrder), string(NotificationErrSync), string(NotificationLowStock), string(NotificationInfo):
		return true
	}
	return false
}

// * MarkAsRead marks the notification as read.
func (notify *Notify) MarkAsRead() {
	notify.ReadStatus = true
}

// * IsRead checks if the notification is read.
func (notify *Notify) IsRead() bool {
	return notify.ReadStatus
}

// * method construnctor
func NewNotification(ownerID int, t NotificationType, msg string) (*Notify, error) {
	notify := &Notify{
		OwnerID:    ownerID,
		Type:       t,
		Message:    msg,
		ReadStatus: false,
		CreatedAt:  time.Now(),
	}
	if err := notify.Validate(); err != nil {
		return nil, err
	}
	return notify, nil
}
