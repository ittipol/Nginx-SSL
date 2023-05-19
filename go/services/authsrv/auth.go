package authsrv

type authResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthService interface {
	Login(email string, password string) (res authResponse, err error)
	Refresh(headers map[string]string) (res authResponse, err error)
	Verify(headers map[string]string) error
}
