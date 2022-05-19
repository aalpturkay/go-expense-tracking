package controllers

import (
	"expenseTracking/models"
	"expenseTracking/services"
	"expenseTracking/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface {
	CreateUserHandler(c *gin.Context)
	FindUserHandler(c *gin.Context)
}

type UserController struct {
	s services.IUserService
}

func NewUserController(s services.IUserService) IUserController {
	return &UserController{s}
}

func (u UserController) CreateUserHandler(c *gin.Context) {
	dto := models.CreateUserDto{}

	if err := c.BindJSON(&dto); err != nil {
		//c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validators.ValidateErrors(err)})
		return
	}

	user := models.User{Name: dto.Name, Balance: 0, Email: dto.Email, Password: dto.Password, Expenses: []models.Expense{}}

	u.s.CreateUser(user)

	c.JSON(http.StatusCreated, &dto)
}

func (u UserController) FindUserHandler(c *gin.Context) {
	dto := models.GetUserDto{}
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	dto, err = u.s.FindUser(int(userId))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, &dto)
}
