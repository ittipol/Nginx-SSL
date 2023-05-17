package repositories

import (
	"database/sql"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (obj userRepository) GetUserByEmail(email string) (user User, err error) {
	// obj.db.Find(&user)

	return User{}, nil
}

func (obj userRepository) GetUserByRefreshToken(id int, refreshToken string) (user User, err error) {
	return User{}, nil
}

func (obj userRepository) SaveRefreshToken(id int, refreshToken string) error {
	return obj.db.Model(&User{}).Where("id = @id", sql.Named("id", id)).Update("refresh_token", refreshToken).Error
}

func (obj userRepository) CreateUser(email string, hashedPassword string, name string) (id int, err error) {

	user := User{
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	}

	tx := obj.db.Create(&user)

	return id, tx.Error
}
