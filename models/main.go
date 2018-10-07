package models

// User stores info required to authentication
type User struct {
	Login    string `json:"login"`
	PassHash string `json:"hash"`
}

// Vacancy stores info about vacancy
type Vacancy struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	Salary     int    `json:"salary"`
	Experience string `json:"experience"`
	Place      string `json:"place"`
}
