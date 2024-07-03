package models

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TrxDebit  = "debit"
	TrxCredit = "credit"

	TrxTypeTopUp   = "top_up"
	TrxTypePayment = "payment"

	TrxStatusSuccess = "success"
	TrxStatusFailed  = "failed"
)

var (
	TrxTypeMap = map[string]string{
		TrxTypeTopUp:   TrxDebit,
		TrxTypePayment: TrxCredit,
	}
)

type TransactionHistory struct {
	TransactionID   uuid.UUID `json:"transaction_id" gorm:"primaryKey"`
	UserID          uuid.UUID `json:"user_id" gorm:"not null"`
	Status          string    `json:"status" gorm:"not null"`
	StatusRemarks   *string   `json:"status_remarks" gorm:"null"`
	Remarks         *string   `json:"remarks" gorm:"null"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`
	Amount          float64   `json:"amount" gorm:"not null"`
	BalanceBefore   float64   `json:"balance_before" gorm:"not null"`
	BalanceAfter    float64   `json:"balance_after" gorm:"not null"`
	CreatedDate     time.Time `json:"created_date,omitempty" gorm:"autoCreateTime"`
}

func (TransactionHistory) TableName() string {
	return "transaction_history"
}

func (u *TransactionHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == uuid.Nil {
		return errors.New("user id is required")
	}

	id, err := uuid.NewV7()
	if err != nil {
		log.Printf("error generate uuid: %v ", err)
		return errors.New("failed generate uuid")
	}
	u.TransactionID = id

	statusRemarks := ""
	status := TrxStatusFailed
	trxType := TrxTypeMap[u.TransactionType]
	switch trxType {
	case TrxDebit:
		u.BalanceAfter = u.BalanceBefore + u.Amount
		status = TrxStatusSuccess
		statusRemarks = u.TransactionType + " success"

	case TrxCredit:
		balanceAfter := u.BalanceBefore - u.Amount
		statusRemarks = "balance not enough"
		if balanceAfter >= 0 {
			statusRemarks = u.TransactionType + " success"
			u.BalanceAfter = balanceAfter
			status = TrxStatusSuccess
		}

	default:
		status = TrxStatusFailed
		statusRemarks = u.TransactionType + " invalid type"
	}

	u.Status = status
	u.StatusRemarks = &statusRemarks

	return
}
