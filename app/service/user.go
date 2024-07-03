package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/adamnasrudin03/go-test-mnc/app/configs"
	"github.com/adamnasrudin03/go-test-mnc/app/dto"
	"github.com/adamnasrudin03/go-test-mnc/app/models"
	"github.com/adamnasrudin03/go-test-mnc/app/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	Register(ctx context.Context, input dto.UserRegisterReq) (res *dto.UserRegisterResp, err *models.RespError)
}

type UserSrv struct {
	Repo   repository.UserRepository
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

// NewUserService creates a new instance of UserService.
func NewUserService(
	UserRepo repository.UserRepository,
	cfg *configs.Configs,
	logger *logrus.Logger,
) UserService {
	return &UserSrv{
		Repo:   UserRepo,
		Cfg:    cfg,
		Logger: logger,
	}
}

func (srv *UserSrv) Register(ctx context.Context, input dto.UserRegisterReq) (res *dto.UserRegisterResp, errResp *models.RespError) {
	const opName = "UserService-Register"

	defer func() {
		if r := recover(); r != nil {
			srv.Logger.Errorf("%v error recover: %v", opName, r)
			res = nil
			errResp = &models.RespError{
				StatusCode: http.StatusInternalServerError,
				Error:      errors.New("internal server error"),
			}
			return
		}

	}()

	err := input.Validate()
	if err != nil {
		errResp = &models.RespError{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
		return nil, errResp
	}

	user := &models.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Pin:         input.Pin,
	}

	err = srv.Repo.CheckDuplicate(ctx, *user)
	if err != nil {
		errResp = &models.RespError{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
		srv.Logger.Errorf("%v error check duplicate: %v", opName, err)
		return nil, errResp
	}

	user, err = srv.Repo.Register(ctx, *user)
	if err != nil {
		errResp = &models.RespError{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
		srv.Logger.Errorf("%v error create data: %v", opName, err)
		return nil, errResp
	}
	if user == nil {
		errResp = &models.RespError{
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("failed create data"),
		}
		return nil, errResp
	}

	res = user.ConvertToResponse()
	return res, nil
}
