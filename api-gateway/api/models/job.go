package models

import "time"

type (
	Job struct {
		ID             string  `json:"id"`
		Name           string  `json:"name"`
		Salary         float32 `json:"salary"`
		Level          string  `json:"level"`
		LocationType   string  `json:"location_type"`
		EmploymentType string  `json:"employment_type"`
		Address        string  `json:"address"`
		Company        string  `json:"company"`
	}

	ResponseJob struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		Salary         float32   `json:"salary"`
		Level          string    `json:"level"`
		LocationType   string    `json:"location_type"`
		EmploymentType string    `json:"employment_type"`
		Address        string    `json:"address"`
		Company        string    `json:"company"`
		StartDate      time.Time `json:"start_date"`
		EndDate        time.Time `json:"end_date"`
	}

	ClientJobs struct {
		ClientID  string `json:"client_id"`
		JobID     string `json:"job_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	ClientJobRequest struct {
		ClientID string `json:"client_id"`
		JobID    string `json:"job_id"`
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
	}

	ClientWithJobs struct {
		Client Client        `json:"client"`
		Jobs   []ResponseJob `json:"jobs"`
	}

	JobWithClients struct {
		Job     ResponseJob `json:"job"`
		Clients []Client    `json:"clients"`
	}
)
