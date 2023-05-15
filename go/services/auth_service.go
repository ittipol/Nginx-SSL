package services

import (
	"go-nginx-ssl/errs"
	"go-nginx-ssl/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var hmacSampleSecret []byte

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{userRepository}
}

func (obj authService) Login(email string, password string) (res authResponse, err error) {

	user, err := obj.userRepository.GetUserByEmail(email)

	if err != nil {
		return res, errs.NewNotFoundError("Invalid username or password")
	}

	// check password matching
	_ = user.Password

	accessToken, _, err := getToken(user.ID)

	if err != nil {
		return res, errs.NewUnexpectedError()
	}

	res = authResponse{
		AccessToken:  accessToken,
		RefreshToken: "",
	}

	return
}

func getToken(id int) (accessToken string, refreshToken string, err error) {

	// Gen access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err = token.SignedString(hmacSampleSecret)

	// Gen refresh token
	//

	return
}
