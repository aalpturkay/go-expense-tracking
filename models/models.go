package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Balance  float64   `json:"balance"`
	Expenses []Expense `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Expense struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	UserID      uint
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type GetUserDto struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email" binding:"required,email"`
	Balance  float64   `json:"balance"`
	Expenses []Expense `json:"expenses"`
}

type CreateExpenseDto struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Category    string  `json:"category" binding:"required,Enum=INCOME_BILLS_TRANSPORT_FOOD"`
	UserID      uint    `json:"user_id" binding:"required"`
}

type UpdateExpenseDto struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category" binding:"Enum=INCOME_BILLS_TRANSPORT_FOOD"`
}
