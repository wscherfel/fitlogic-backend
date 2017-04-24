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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":time.Now().Add(fitlogic.JWTExpiration).Unix(),
		"userId": userId,
		"role": role,
	})

	tokenString, err := token.SignedString([]byte(fitlogic.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIdAndRoleFromToken(token *jwt.Token) (id uint, role int, err error) {
	claims := token.Claims.(jwt.MapClaims)

	nonTypedId, foundId := claims["userId"]
	floatId := nonTypedId.(float64)
	id = uint(floatId)

	nonTypedRole, foundRole := claims["role"]
	floatRole := nonTypedRole.(float64)
	role = int(floatRole)

	if !foundId || !foundRole {
		return 0, 0, ErrMissingTokenClaims
	}

	return id, role, nil
}
