package models

import "time"

type Appointment struct {
	ID 			int 		`json:"id,omitempty"`
	Brand 		string 		`json:"brand,omitempty"`
	Model 		string 		`json:"model,omitempty"`
	DateJoined 	time.Time 	`json:"datejoined,omitempty"`
	Owner 		*Person 	`json:"owner,omitempty"`
}

var Appointments []*Appointment

