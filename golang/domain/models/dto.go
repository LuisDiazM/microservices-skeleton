package models

import "time"

type UserResponsesDTO struct {
	Answer   string    `json:"answer,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	Contact  string    `json:"contact,omitempty"`
}
