package db

import (
	"expenseTracking/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// docker run --name <CONTAINER_ADI> -e POSTGRES_PASSWORD=<ROOT_PAROLASI> -d -p 5432:5432 postgres

func ConnectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=alpthedev dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Expense{})
	if err != nil {
		panic(err)
	}
	log.Println("Veri tabanı bağlantısı başarılı!")

	return db
}
