package databases

import (
	"TiBO_API/repository/databases/addressesRepo"
	"TiBO_API/repository/databases/cinemasRepo"
	"TiBO_API/repository/databases/invoiceRepo"
	"TiBO_API/repository/databases/moviesRepo"
	"TiBO_API/repository/databases/usersRepo"
	"fmt"
	"os"
	"runtime"

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

	var dbHost string
	if runtime.GOOS != "windows" {
		dbHost = os.Getenv("DB_HOST_DOCKER")
	}else {
		dbHost = os.Getenv("DB_HOST")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
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
	db.AutoMigrate(&addressesRepo.Addresses{})
	db.AutoMigrate(&cinemasRepo.Cinemas{})
	db.AutoMigrate(&moviesRepo.Movies{})
	db.AutoMigrate(&invoiceRepo.Invoices{})
}
