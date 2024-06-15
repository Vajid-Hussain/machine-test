package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Conver password to hashed string
func HashPassword(password string) (string, error) {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("fase issue during password hashing ")
	}
	return string(HashedPassword), nil
}

// Compare plain password with hashed one
func CompairPassword(hashedPassword string, plainPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("password does't match")
	}

	return nil
}

// Generate access token
func GenerateAcessToken(securityKey string, id string, duration int64) (string, error) {

	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + duration,
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

// Generate refresh token
func GenerateRefreshToken(securityKey, id string) (string, error) {
	key := []byte(securityKey)
	clamis := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Unix() + 360000,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("making refresh token lead to error")
	}

	return signedToken, nil
}

// Verify tokens
func VerifyAcessToken(token string, secretkey string) (string, error) {

	key := []byte(secretkey)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if parsedToken == nil {
		return "", errors.New("invalid access token")
	}

	if err != nil {
		return "", err
	}

	if len(parsedToken.Header) == 0 {
		return "", errors.New("token tamberd include header")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(string)

	if !ok {
		return "", errors.New("id is not in accessToken. access denied")
	}
	return id, nil
}

// Struct validation
func Validator(request any) []string {
	errResponse := []string{}

	var validate = validator.New()
	errs := validate.Struct(request)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		errResponse = append(errResponse, fmt.Sprintf("[%s] : '%v' become '%s' %s", err.Field(), err.Value(), err.Tag(), err.Param()))
	}

	return errResponse
}

// Pagination
func Pagination(limit, offset string) (string, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return "", err
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "", err
	}

	if limitInt < 1 || offsetInt < 1 {
		return "", responsemodels.ErrPaginationWrongValue
	}

	return strconv.Itoa((offsetInt * limitInt) - limitInt), nil
}
