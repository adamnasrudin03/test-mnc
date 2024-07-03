package main

import (
	"fmt"
	"log"

	"github.com/adamnasrudin03/go-test-mnc/app"
	"github.com/adamnasrudin03/go-test-mnc/app/configs"
	"github.com/adamnasrudin03/go-test-mnc/app/router"
	"github.com/adamnasrudin03/go-test-mnc/pkg/database"
	"github.com/adamnasrudin03/go-test-mnc/pkg/driver"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load env file")
	}

	var (
		cfg                  = configs.GetInstance()
		logger               = driver.Logger(cfg)
		db          *gorm.DB = database.SetupDbConnection(cfg, logger)
		repo                 = app.WiringRepository(db, cfg, logger)
		services             = app.WiringService(repo, cfg, logger)
		controllers          = app.WiringController(services, cfg, logger)
	)

	defer database.CloseDbConnection(db, logger)

	r := router.NewRoutes(*controllers)

	listen := fmt.Sprintf(":%v", cfg.App.Port)
	r.Run(listen)
}
