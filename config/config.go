package config

import (
	"github.com/NonsoAmadi10/bitcoin-wallet/utils"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB(model ...interface{}) {

	dbString := utils.GetEnv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect to Database")
	}

	db.AutoMigrate(model...)

	DB = db
}
