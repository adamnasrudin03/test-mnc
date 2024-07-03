package models

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	WalletID    uuid.UUID `json:"wallet_id" gorm:"primaryKey"`
	UserID      uuid.UUID `json:"user_id" gorm:"not null;unique"`
	Amount      float64   `json:"amount" gorm:"not null"`
	CreatedDate time.Time `json:"created_date,omitempty" gorm:"autoCreateTime"`
	UpdatedDate time.Time `json:"updated_date,omitempty" gorm:"autoUpdateTime"`
}

func (Wallet) TableName() string {
	return "wallet"
}
func (u *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == uuid.Nil {
		return errors.New("user id is required")
	}

	id, err := uuid.NewV7()
	if err != nil {
		log.Printf("error generate uuid: %v ", err)
		return errors.New("failed generate uuid")
	}
	u.WalletID = id

	return
}
