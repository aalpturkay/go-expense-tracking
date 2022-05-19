package controllers

import (
	"expenseTracking/models"
	"expenseTracking/services"
	"expenseTracking/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IExpenseController interface {
	CreateExpenseHandler(c *gin.Context)
	DeleteExpenseHandler(c *gin.Context)
	UpdateExpenseHandler(c *gin.Context)
}

type ExpenseController struct {
	s services.IExpenseService
}

func NewExpenseController(s services.IExpenseService) IExpenseController {
	return &ExpenseController{s}
}

func (u ExpenseController) CreateExpenseHandler(c *gin.Context) {
	dto := models.CreateExpenseDto{}

	if err := c.BindJSON(&dto); err != nil {
		//c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validators.ValidateErrors(err)})
		return
	}

	expense := models.Expense{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		Category:    dto.Category,
		UserID:      dto.UserID,
	}

	u.s.CreateExpense(expense)

	c.JSON(http.StatusCreated, &dto)
}

func (u ExpenseController) DeleteExpenseHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	if err := u.s.DeleteExpense(int(id)); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.Status(http.StatusOK)
}

func (u ExpenseController) UpdateExpenseHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var dto models.UpdateExpenseDto
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validators.ValidateErrors(err)})
		return
	}

	if err := u.s.UpdateExpense(dto, int(id)); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Status(http.StatusOK)
}
