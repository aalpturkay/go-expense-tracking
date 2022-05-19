package main

import (
	"expenseTracking/controllers"
	"expenseTracking/db"
	"expenseTracking/repositories"
	"expenseTracking/routes"
	"expenseTracking/services"
	"expenseTracking/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	dbInstance := db.ConnectDb()
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Enum", validators.Enum)
	}

	expenseRepository := repositories.NewExpenseRepository(dbInstance)
	expenseService := services.NewExpenseService(expenseRepository)
	expenseController := controllers.NewExpenseController(expenseService)

	userRepository := repositories.NewUserRepository(dbInstance)
	userService := services.NewUserService(userRepository, expenseService)
	userController := controllers.NewUserController(userService)

	routes.InitializeUsersRoutes(r, userController)
	routes.InitializeExpensesRoutes(r, expenseController)

	r.Run(":4435")
}

/* Eksliklikler:
- Auth özelliği (JWT Token)
- Parolanın bir hash fonksiyonu kullanılarak şifrelenmesi
- Kullanıcıların sadece izin verilen endpointlere erişim sağlayabilmesi
- Harcamanın filtrelenmesi
- Kod için, test edilebilir ve daha az Dependency isteyen yapılar.
*/
