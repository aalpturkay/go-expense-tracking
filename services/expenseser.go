package services

import (
	"expenseTracking/models"
	"expenseTracking/repositories"
)

type IExpenseService interface {
	CreateExpense(expense models.Expense)
	GetExpenses(userId int) ([]models.Expense, error)
	DeleteExpense(expenseId int) error
	UpdateExpense(dto models.UpdateExpenseDto, expenseId int) error
}

type ExpenseService struct {
	r repositories.IExpenseRepository
}

func NewExpenseService(r repositories.IExpenseRepository) IExpenseService {
	return &ExpenseService{r}
}

func (s ExpenseService) CreateExpense(expense models.Expense) {
	s.r.CreateExpense(expense)
}

func (s ExpenseService) GetExpenses(userId int) ([]models.Expense, error) {
	expenses, err := s.r.GetExpenses(userId)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s ExpenseService) DeleteExpense(expenseId int) error {
	err := s.r.DeleteExpense(expenseId)
	if err != nil {
		return err
	}
	return nil
}

func (s ExpenseService) UpdateExpense(dto models.UpdateExpenseDto, expenseId int) error {
	err := s.r.UpdateExpense(dto, expenseId)
	if err != nil {
		return err
	}
	return nil
}
