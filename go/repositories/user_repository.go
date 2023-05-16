package repositories

import "gorm.io/gorm"

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
