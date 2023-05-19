package usersrv

import (
	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/errs"
	"go-nginx-ssl/helpers"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repositories.UserRepository
	jwtToken       appUtils.JwtUtil
}

func NewUserService(userRepository repositories.UserRepository, jwtToken appUtils.JwtUtil) UserService {
	return &userService{userRepository, jwtToken}
}

func (obj userService) Register(email string, password string, name string) error {

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		logs.Error(err)
		return errs.NewUnauthorizedError()
	}

	_, err = obj.userRepository.CreateUser(email, string(hashedPassword), name)

	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (obj userService) Profile(headers map[string]string) (res profileResponse, err error) {

	value, _ := helpers.GetHeader(headers, "Authorization")

	tokenString, _ := helpers.GetBearerToken(value)

	token, _ := obj.jwtToken.Validate(tokenString, appUtils.AccessTokenSecretKey)

	claims, _ := token.Claims.(jwt.MapClaims)

	id := int(claims["id"].(float64))

	user, err := obj.userRepository.GetUserById(id)

	if err != nil {
		return res, errs.NewNotFoundError("User not found")
	}

	res.Name = user.Name

	return res, err
}
