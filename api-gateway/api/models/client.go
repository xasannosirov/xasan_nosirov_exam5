package models

type (
	Client struct {
		Id          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Age         uint64 `json:"age"`
		Gender      string `json:"gender"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		Status      bool   `json:"status"`
		Refresh     string `json:"refresh"`
	}

	Status struct {
		Status bool `json:"status"`
	}
)
