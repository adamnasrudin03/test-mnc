package models

import (
	"errors"
	"log"
	"time"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/adamnasrudin03/go-test-mnc/app/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID      uuid.UUID `json:"user_id" gorm:"primaryKey"`
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	Address     string    `json:"address" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null;unique"`
	Pin         string    `json:"pin,omitempty" gorm:"not null"`
	CreatedDate time.Time `json:"created_date,omitempty" gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "user"
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPass, err := helpers.HashPassword(u.Pin)
	if err != nil {
		log.Printf("failed hash Pin: %v ", err)
		return errors.New("failed hash Pin")
	}

	u.Pin = hashedPass
	id, err := uuid.NewV7()
	if err != nil {
		log.Printf("error generate uuid: %v ", err)
		return errors.New("failed generate uuid")
	}
	u.UserID = id

	return
}

func (u *User) ConvertToResponse() *dto.UserRegisterResp {
	return &dto.UserRegisterResp{
		UserID:      u.UserID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		CreatedDate: u.CreatedDate.Format(helpers.FormatDateTime),
	}
}
