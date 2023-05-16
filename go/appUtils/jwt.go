package appUtils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	accessTokenSecret  []byte = []byte(viper.GetString("jwtAccessTokenSecret"))
	refreshTokenSecret []byte = []byte(viper.GetString("jwtRefreshTokenSecret"))
)

type accessTokenClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type refreshTokenClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type JwtUtil interface {
	GenToken(id int) (accessToken string, refreshToken string, err error)
	Validate(tokenString string) (token *jwt.Token, err error)
}

type jwtUtil struct {
}

func NewJwtUtil() JwtUtil {
	return &jwtUtil{}
}

func (obj jwtUtil) GenToken(id int) (accessToken string, refreshToken string, err error) {

	// accessTokenSecret = []byte(viper.GetString("jwtAccessTokenSecret"))
	// refreshTokenSecret = []byte(viper.GetString("jwtRefreshTokenSecret"))

	// Gen access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims{
		id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(400 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the secret
	accessToken, err = token.SignedString(accessTokenSecret)

	if err != nil {
		return
	}

	// Gen refresh token
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims{
		id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the secret
	refreshToken, err = token.SignedString(refreshTokenSecret)

	return
}

func (obj jwtUtil) Validate(tokenString string) (token *jwt.Token, err error) {

	// accessTokenSecret = []byte(viper.GetString("jwtAccessTokenSecret"))

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return accessTokenSecret, nil
	})

	return
}
