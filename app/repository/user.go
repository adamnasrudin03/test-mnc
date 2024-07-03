package repository

import (
	"context"
	"errors"

	"github.com/adamnasrudin03/go-test-mnc/app/configs"
	"github.com/adamnasrudin03/go-test-mnc/app/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CheckDuplicate(ctx context.Context, input models.User) (err error)
	Register(ctx context.Context, input models.User) (res *models.User, err error)
}

type UserRepo struct {
	DB     *gorm.DB
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewUserRepository(
	db *gorm.DB,
	cfg *configs.Configs,
	logger *logrus.Logger,
) UserRepository {
	return &UserRepo{
		DB:     db,
		Cfg:    cfg,
		Logger: logger,
	}
}

func (r *UserRepo) trxEnd(trx *gorm.DB, err error) {
	if rc := recover(); rc != nil {
		r.Logger.Errorf(`trxEnd Panic Error %v`, r)
		trx.Rollback()
		return
	}
	if err != nil {
		r.Logger.Errorf(`trxEnd Error %v`, err)
		trx.Rollback()
		return
	}
	if err := trx.Commit(); err != nil {
		r.Logger.Errorf(`trxEnd err commit %v`, err)
		trx.Rollback()
		return
	}
}

func (r *UserRepo) CheckDuplicate(ctx context.Context, input models.User) (err error) {
	user := new(models.User)
	_ = r.DB.WithContext(ctx).Select("user_id").Where("phone_number = ?", input.PhoneNumber).First(&user).Error
	if user != nil && user.UserID != uuid.Nil {
		return errors.New("phone number already registered")
	}

	return nil
}

func (r *UserRepo) Register(ctx context.Context, input models.User) (res *models.User, err error) {
	const opName = "UserRepository-Register"

	trx := r.DB.Begin().WithContext(ctx)
	defer func() {
		r.trxEnd(trx, err)
	}()

	err = trx.Clauses(clause.Returning{}).Create(&input).Error
	if err != nil {
		r.Logger.Errorf("%v error register new user: %v ", opName, err)
		return nil, errors.New("failed register new user")
	}

	wallet := models.Wallet{
		UserID: input.UserID,
		Amount: 0,
	}
	err = trx.Clauses(clause.Returning{}).Create(&wallet).Error
	if err != nil {
		r.Logger.Errorf("%v error register new wallet: %v ", opName, err)
		return nil, errors.New("failed create new wallet")
	}

	return &input, nil
}
