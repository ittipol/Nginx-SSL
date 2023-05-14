package services

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (obj authService) Login(email string, password string) (res authResponse, err error) {

	res = authResponse{
		AccessToken:  "",
		RefreshToken: "",
	}

	return
}
