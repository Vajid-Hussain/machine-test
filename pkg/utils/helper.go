package utils

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

func CreateSession(cfg config.S3Bucket) *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
			Endpoint: aws.String(""),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}

func UploadImageToS3(file []byte, sess *session.Session) (string, error) {

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("hyper-hive-data"),
		Key:    aws.String("chat media/" + fileName),
		Body:   aws.ReadSeekCloser(bytes.NewReader(file)),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println("err from s3 upload", err)
		return "", err
	}
	return upload.Location, nil
}
