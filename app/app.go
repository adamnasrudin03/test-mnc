package app

import (
	"github.com/adamnasrudin03/go-test-mnc/app/configs"
	"github.com/adamnasrudin03/go-test-mnc/app/controller"
	"github.com/adamnasrudin03/go-test-mnc/app/repository"
	"github.com/adamnasrudin03/go-test-mnc/app/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB, cfg *configs.Configs, logger *logrus.Logger) *repository.Repositories {
	return &repository.Repositories{
		User: repository.NewUserRepository(db, cfg, logger),
	}
}

func WiringService(repo *repository.Repositories, cfg *configs.Configs, logger *logrus.Logger) *service.Services {
	return &service.Services{
		User: service.NewUserService(repo.User, cfg, logger),
	}
}

func WiringController(srv *service.Services, cfg *configs.Configs, logger *logrus.Logger) *controller.Controllers {
	return &controller.Controllers{
		User: controller.NewUserDelivery(srv.User, logger),
	}
}
