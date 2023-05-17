package repositories

type User struct {
	ID           int
	Password     string
	Email        string `gorm:"unique:size:100"`
	Name         string `gorm:"size:100"`
	RefreshToken string
}

type UserRepository interface {
	GetUserByEmail(email string) (user User, err error)
	GetUserByRefreshToken(id int, refreshToken string) (user User, err error)
	SaveRefreshToken(id int, refreshToken string) error
	CreateUser(email string, hashedPassword string, name string) (id int, err error)
}
