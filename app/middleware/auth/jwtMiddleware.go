package auth

import (
	"TiBO_API/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Uuid string `json:"uuid"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT   string
	ExpDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden,
				helpers.BuildErrorResponse("failed to init token",
					e, helpers.EmptyObj{}))
		}),
	}
}

// GenerateToken jwt
func (cj *ConfigJWT) GenerateToken(userID string, UserRole string) string {
	claims := JwtCustomClaims{
		userID,
		UserRole,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpDuration))).Unix(),
		},
	}
	// CreateCalorie token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(cj.SecretJWT))

	return token
}

//get user
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
