package usersrv

type profileResponse struct {
	Name string `json:"name"`
}

type UserService interface {
	Register(email string, password string, name string) error
	Profile(headers map[string]string) (res profileResponse, err error)
}
