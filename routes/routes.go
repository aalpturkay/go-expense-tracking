package routes

import (
	"expenseTracking/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUsersRoutes(r *gin.Engine, controller controllers.IUserController) {
	routes := r.Group("/users")
	routes.POST("/", controller.CreateUserHandler)
	routes.GET("/:id", controller.FindUserHandler)
}

func InitializeExpensesRoutes(r *gin.Engine, controller controllers.IExpenseController) {
	routes := r.Group("/expenses")
	routes.POST("/", controller.CreateExpenseHandler)
	routes.PUT("/:id", controller.UpdateExpenseHandler)
	routes.DELETE("/:id", controller.DeleteExpenseHandler)
}
