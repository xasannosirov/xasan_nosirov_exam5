package entity

import "time"

type Job struct {
	GUID           string
	Name           string
	Salary         float64
	Level          string
	LocationType   string
	EmploymentType string
	Address        string
	Company        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ClientJob struct {
	ClientID  string
	JobID     string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobClientRequest struct {
	JobID    string
	ClientID string
	Page     uint64
	Limit    uint64
}

type Response struct {
	Status bool
}
