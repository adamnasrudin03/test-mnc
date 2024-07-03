package controller

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/adamnasrudin03/go-test-mnc/app/dto"
	"github.com/adamnasrudin03/go-test-mnc/app/service"
	"github.com/adamnasrudin03/go-test-mnc/pkg/helpers/response_mapper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController interface {
	Register(ctx *gin.Context)
}

type UserHandler struct {
	Service service.UserService
	Logger  *logrus.Logger
}

func NewUserDelivery(
	srv service.UserService,
	logger *logrus.Logger,
) UserController {
	return &UserHandler{
		Service: srv,
		Logger:  logger,
	}
}

func (c *UserHandler) Register(ctx *gin.Context) {
	var (
		opName = "UserController-Register"
		input  dto.UserRegisterReq
		err    error
	)

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	res, errResp := c.Service.Register(ctx, input)
	if errResp != nil {
		c.Logger.Errorf("%v error: %v ", opName, errResp)
		ctx.JSON(errResp.StatusCode, response_mapper.APIResponse(errResp.Error.Error(), errResp.StatusCode, nil))
		return
	}

	ctx.JSON(http.StatusCreated, response_mapper.APIResponse("", http.StatusCreated, res))
}
