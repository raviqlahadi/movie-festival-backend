package repositories

import (
	"github.com/raviqlahadi/movie-festival-backend/internal/db"
	"github.com/raviqlahadi/movie-festival-backend/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user models.User) error {
	return db.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
