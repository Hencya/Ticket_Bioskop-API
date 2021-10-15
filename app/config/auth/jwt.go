package auth

import (
	"TiBO_API/app/middleware/auth"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func SetupJwt() auth.ConfigJWT {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	jwtScret := os.Getenv("JWT_SECRET")
	jwtExp := os.Getenv("JWT_EXPIRE")

	Exp, _ := strconv.Atoi(jwtExp)
	configJWT := auth.ConfigJWT{
		SecretJWT:   jwtScret,
		ExpDuration: Exp,
	}

	return configJWT
}
