package models

import (
	"errors"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

/*
* Client Represents a client in the e-commerce application.
* It contains fields for the client ID, first name, last name, phone number, email
* address, password
* This struct is used to store information about clients in the database.
 */
type Client struct {
	User
	Age       int `json:"age"`
	AddressID int `json:"address_id"`
}

// * methods for Client struct

// Valiidate checks if the Client has all the required fields to be created or updated.
func (client *Client) Validate() error {
	if strings.TrimSpace(client.FirstName) == "" {
		return errors.New("first name is required")
	}
	if strings.TrimSpace(client.LastName) == "" {
		return errors.New("last name is required")
	}
	if err := client.IsValidPhone(client.Phone); err != nil {
		return err
	}
	if strings.TrimSpace(client.Email) == "" {
		return errors.New("email address is required")
	}
	if strings.TrimSpace(client.Password) == "" {
		return errors.New("password is required")
	}
	if client.Age <= 0 {
		return errors.New("age must be greater than 0")
	}

	// validate if gender is male or female
	if !client.IsValidGender(client.Gender) {
		return errors.New("gender must be male, female, or other")
	}

	return nil
}

// IsValidGender is a validator of the gender field.
func (client *Client) IsValidGender(gender string) bool {
	switch strings.ToLower(gender) {
	case "male", "female", "other":
		return true
	}
	return false
}

// IsValidPhone checks if the phone number is valid.
func (client *Client) IsValidPhone(phone string) error {
	if len(phone) <= 8 {
		return errors.New("phone number must be at least 8 characters long")
	}
	return nil
}

// IsValidEmail checks if the email address is valid.
func (client *Client) IsValidEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email address is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email address")
	}

	return nil
}

// VerifyPassword checks if the password is valid.
func (client *Client) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	return err == nil
}

// HashPassword hashes the password using bcrypt.
func (client *Client) HashPassword(password string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	client.Password = string(hashed_password)
	return nil
}

func (client *Client) FullName() string {
	return client.FirstName + " " + client.LastName
}
