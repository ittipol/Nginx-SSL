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

func (obj userRepository) GetUserById(id int) (user User, err error) {
	err = obj.db.Where("id = @id", sql.Named("id", id)).First(&user).Error
	return user, err
}

func (obj userRepository) GetUserByEmail(email string) (user User, err error) {
	err = obj.db.Where("email = @email", sql.Named("email", email)).First(&user).Error
	return user, err
}

func (obj userRepository) GetUserByRefreshToken(id int, refreshToken string) (user User, err error) {
	err = obj.db.Where(&User{ID: id, RefreshToken: refreshToken}).First(&user).Error
	return user, err
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

	err = obj.db.Create(&user).Error

	return id, err
}
