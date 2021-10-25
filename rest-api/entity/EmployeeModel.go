package entity

import "time"

type Employee struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"firstname"`
	MiddleName   string    `json:"middlename"`
	LastName     string    `json:"lastname"`
	Gender       string    `json:"gender"`
	Salary       float64   `json:"salary"`
	DOB          time.Time `json:"dob"`
	Email        string    `json:"email`
	Phone        int       `json:"phone"`
	AddressLine1 string    `json:"address1"`
	AddressLine2 string    `json:"address2"`
	State        string    `json:"state"`
	PostCode     int       `json:"postcode"`
	TFN          int       `json:"tfn"`
	SuperBalance float64   `json:"superbalance"`
}
