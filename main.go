package main

import (
	"fmt"
	"go-crud/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=dev dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("Error connecting to database")
	}
	fmt.Println("Database " + db.Name() + " connected at port: 5432")
	migration := models.CreateTables(db)
	testUser := models.User{ID: "1!", Name: "Test name", Email: "test@test.test"}
	db.Create(&testUser)
	if migration != nil {
		fmt.Println("Error pushing models into database")
	} else {
		fmt.Println("Models succesfully pushed into database")
	}

}
