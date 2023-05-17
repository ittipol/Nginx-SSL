package usersrv

import (
	"go-nginx-ssl/errs"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/repositories"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
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
