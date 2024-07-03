package dto

import (
	"strings"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type UserRegisterReq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin,omitempty"`
}

func (u *UserRegisterReq) Validate() (err error) {
	u.FirstName = strings.TrimSpace(u.FirstName)
	if u.FirstName == "" {
		return helpers.ErrIsRequired("nama depan", "first name")
	}

	u.LastName = strings.TrimSpace(u.LastName)
	u.Address = strings.TrimSpace(u.Address)
	if u.Address == "" {
		return helpers.ErrIsRequired("alamat", "address")
	}

	u.PhoneNumber = strings.TrimSpace(u.PhoneNumber)
	if u.PhoneNumber == "" {
		return helpers.ErrIsRequired("nomor telepon", "phone number")
	}

	u.Pin = strings.TrimSpace(u.Pin)
	if u.Pin == "" {
		return helpers.ErrIsRequired("pin", "pin")
	}

	return nil
}
