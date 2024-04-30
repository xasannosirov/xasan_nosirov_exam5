package entity

import "time"

type Client struct {
	GUID        string
	FirstName   string
	LastName    string
	Age         uint64
	Gender      string
	PhoneNumber string
	Address     string
	Email       string
	Password    string
	Status      bool
	Refresh     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type IsUnique struct {
	Email string
}

type UpdateRefresh struct {
	ClientID     string
	RefreshToken string
}

type UpdatePassword struct {
	ClientID    string
	NewPassword string
}

type Response struct {
	Status bool
}
