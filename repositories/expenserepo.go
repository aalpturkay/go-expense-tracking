package repositories

import (
	"expenseTracking/models"
	"gorm.io/gorm"
	"log"
)

type IExpenseRepository interface {
	CreateExpense(expense models.Expense)
	GetExpenses(userId int) ([]models.Expense, error)
	DeleteExpense(expenseId int) error
	UpdateExpense(dto models.UpdateExpenseDto, expenseId int) error
}

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) IExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (u *ExpenseRepository) CreateExpense(expense models.Expense) {
	var user models.User
	u.db.First(&user, expense.UserID)
	log.Println(expense.Category)
	if expense.Category == "INCOME" {
		user.Balance += expense.Price
	} else {
		user.Balance -= expense.Price
	}
	u.db.Save(&user)
	u.db.Create(&expense)
}

func (u ExpenseRepository) GetExpenses(userId int) ([]models.Expense, error) {
	var expenses []models.Expense
	result := u.db.Where(&models.Expense{UserID: uint(userId)}).Find(&expenses)
	return expenses, result.Error
}

func (u ExpenseRepository) DeleteExpense(expenseId int) error {
	var expense models.Expense
	result := u.db.First(&expense, expenseId)
	if result.Error != nil {
		return result.Error
	}
	u.db.Delete(&expense)
	return nil
}

func (u ExpenseRepository) UpdateExpense(dto models.UpdateExpenseDto, expenseId int) error {
	var expense models.Expense
	result := u.db.First(&expense, expenseId)
	if result.Error != nil {
		return result.Error
	}

	expense.Title = dto.Title
	expense.Description = dto.Description
	expense.Category = dto.Category
	expense.Price = dto.Price

	u.db.Save(&expense)
	return nil
}
