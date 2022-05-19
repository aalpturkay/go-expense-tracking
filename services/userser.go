package services

import (
	"expenseTracking/models"
	"expenseTracking/repositories"
)

type IUserService interface {
	CreateUser(user models.User)
	FindUser(userId int) (models.GetUserDto, error)
}

type UserService struct {
	r              repositories.IUserRepository
	expenseService IExpenseService
}

func NewUserService(r repositories.IUserRepository, expenseSer IExpenseService) IUserService {
	return &UserService{r, expenseSer}
}

func (s UserService) CreateUser(user models.User) {
	s.r.CreateUser(user)
}

func (s UserService) FindUser(userId int) (models.GetUserDto, error) {
	user, err := s.r.FindUser(userId)

	expenses, err := s.expenseService.GetExpenses(userId)

	return models.GetUserDto{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Balance:  user.Balance,
		Expenses: expenses,
	}, err
}
