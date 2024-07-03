package dto

import (
	"github.com/google/uuid"
)

type UserRegisterResp struct {
	UserID      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedDate string    `json:"created_date"`
}
