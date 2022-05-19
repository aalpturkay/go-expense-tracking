package repositories

import (
	"expenseTracking/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user models.User)
	FindUser(userId int) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user models.User) {
	u.db.Create(&user)
}

func (u UserRepository) FindUser(userId int) (models.User, error) {
	var user models.User
	result := u.db.Find(&user, userId)
	return user, result.Error
}
