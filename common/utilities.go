package common

import (
	"github.com/labstack/echo"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/wscherfel/fitlogic-backend"
)


func BindAndValid(c echo.Context, model interface{}) error {
	if err := c.Bind(model); err != nil {
		return err
	}

	valid, err := govalidator.ValidateStruct(model)
	if err != nil && !valid {
		return err
	}

	return nil
}

func CreateToken(userId uint, role int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(fitlogic.JWTExpiration).Unix()
	claims["name"]= userId
	claims["admin"]= role

	tokenString, err := token.SignedString([]byte(fitlogic.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
