package services

import (
	"fmt"

	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/errs"
	"go-nginx-ssl/helpers"
	"go-nginx-ssl/repositories"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	userRepository repositories.UserRepository
	jwtToken       appUtils.JwtUtil
}

func NewAuthService(userRepository repositories.UserRepository, jwtToken appUtils.JwtUtil) AuthService {
	return &authService{userRepository, jwtToken}
}

func (obj authService) Login(email string, password string) (res authResponse, err error) {

	user, err := obj.userRepository.GetUserByEmail(email)

	if err != nil {
		return res, errs.NewNotFoundError("Invalid username or password")
	}

	// check password matching
	_ = user.Password

	accessToken, refreshToken, err := obj.jwtToken.GenToken(user.ID)

	if err != nil {
		return res, errs.NewUnexpectedError()
	}

	res = authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

func (obj authService) Refresh(headers map[string]string) (res authResponse, err error) {

	value, err := helpers.GetHeader(headers, "Authorization")

	if err != nil {
		return res, errs.NewBadRequestError()
	}

	tokenString, err := helpers.GetBearerToken(value)

	if err != nil {
		return res, errs.NewBadRequestError()
	}

	// Check latest refresh token in database
	// user := obj.userRepository.CheckRefreshTokenExist(id, tokenString)

	// check token is valid
	token, err := obj.jwtToken.Validate(tokenString)

	if err != nil {
		return res, errs.NewUnauthorizedError()
	}

	_ = token

	// gen new token
	// accessToken, refreshToken, err := obj.jwtToken.GenToken(user.ID)

	// if err != nil {
	// 	return res, errs.NewUnexpectedError()
	// }

	// res = authResponse{
	// 	AccessToken:  accessToken,
	// 	RefreshToken: refreshToken,
	// }

	return res, nil
}

func (obj authService) Verify(headers map[string]string) error {

	value, err := helpers.GetHeader(headers, "Authorization")

	if err != nil {
		return errs.NewBadRequestError()
	}

	tokenString, err := helpers.GetBearerToken(value)

	if err != nil {
		return errs.NewBadRequestError()
	}

	token, err := obj.jwtToken.Validate(tokenString)

	fmt.Printf("Token Valid: %v \n", token.Valid)

	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
		return errs.NewError(http.StatusOK, "Invalid Token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Res: %v \n", claims)
	}

	return nil
}
