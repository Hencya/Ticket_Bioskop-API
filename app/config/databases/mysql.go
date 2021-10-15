package databases

import (
	"TiBO_API/repository/databases/usersRepo"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	var dbName string
	if os.Getenv("ENV") == "TESTING" {
		dbName = os.Getenv("DB_NAME_TESTING")
	} else {
		dbName = os.Getenv("DB_NAME")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//auto migrate
	dbMigrate(db)

	return db
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&usersRepo.Users{})
}
