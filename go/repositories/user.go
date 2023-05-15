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
}
