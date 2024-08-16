package gormfx

import (
	"github.com/priscila-albertini-silva/jaded-backend/config"
	"github.com/priscila-albertini-silva/jaded-backend/internal/models"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDatabaseConnection() (*gorm.DB, error) {
	host := config.Configuration.Database.Host

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: host,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Could not connect to database: %+v", err)
		return nil, err
	}

	log.Infof("Database connected to: %+v", gormDB)

	gormDB.AutoMigrate(
		&models.Stock{},
		// &models.Dividend{},
		// &models.Order{},
	)

	return gormDB, err
}

var Module = fx.Module("gorm", fx.Invoke(initDatabaseConnection))
