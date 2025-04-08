package database

import (
	"fmt"
	"todo_list/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := "localhost"
	user := "postgres"
	password := "root"
	dbname := "demo"
	port := 5432

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, fmt.Sprintf("%d", port),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	DB.AutoMigrate(&models.Todo{})
	fmt.Println("Connected to the database successfully")
}
