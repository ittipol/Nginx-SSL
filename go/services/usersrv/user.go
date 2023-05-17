package usersrv

type UserService interface {
	Register(email string, password string, name string) error
}
