package common

import (
	"github.com/labstack/echo"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/spf13/viper"
)

const (
	JWTExpiration = time.Hour * 24 * 5
)

var (
	DateMin = time.Date(1970, time.January, 1, 0,0,0,0, time.UTC)

	DateMax = time.Date(2099, time.January, 1, 0,0,0,0, time.UTC)
)

// IDsRequest is used whenever user calls API and sends an array of IDs
type IDsRequest struct{
	IDs []uint `valid:"required"`
}

// BindAndValid will bind body of request to given type and check if its valid
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

// CreateToken will create token with user ID and role coded inside and
// return its string form
func CreateToken(userId uint, role int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":time.Now().Add(JWTExpiration).Unix(),
		"userId": userId,
		"role": role,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("Secret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetUserIdAndRoleFromToken will decode token and return user id and his role
// from token
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
