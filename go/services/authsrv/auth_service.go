package authsrv

import (
	"fmt"

	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/errs"
	"go-nginx-ssl/helpers"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/repositories"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

	// Check password matching
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		logs.Error(err)
		return res, errs.NewUnexpectedError()
	}

	accessToken, refreshToken, err := obj.jwtToken.GenToken(user.ID)

	if err != nil {
		logs.Error(err)
		return res, errs.NewUnexpectedError()
	}

	// Save new refresh token into database
	obj.userRepository.SaveRefreshToken(user.ID, refreshToken)

	res = authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

func (obj authService) Refresh(headers map[string]string) (res authResponse, err error) {

	value, _ := helpers.GetHeader(headers, "Authorization")

	tokenString, _ := helpers.GetBearerToken(value)

	token, _ := obj.jwtToken.Validate(tokenString, appUtils.AccessTokenSecretKey)

	claims, _ := token.Claims.(jwt.MapClaims)

	id := int(claims["id"].(float64))

	// Check it's latest refresh token in database
	user, err := obj.userRepository.GetUserByRefreshToken(id, tokenString)

	if err != nil {
		logs.Error(err)
		return res, errs.NewUnexpectedError()
	}

	// Gen new token
	accessToken, refreshToken, err := obj.jwtToken.GenToken(user.ID)

	if err != nil {
		return res, errs.NewUnexpectedError()
	}

	// Save new refresh token into database
	obj.userRepository.SaveRefreshToken(user.ID, refreshToken)

	res = authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

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

	token, err := obj.jwtToken.Validate(tokenString, appUtils.AccessTokenSecretKey)

	fmt.Printf("Token Valid: %v \n", token.Valid)

	if err != nil {
		logs.Error(fmt.Sprintf("Error: %v \n", err.Error()))
		return errs.NewError(http.StatusOK, "Invalid Token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Res: %v \n", claims)
	}

	return nil
}
