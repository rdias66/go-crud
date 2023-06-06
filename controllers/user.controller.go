package users_controller

import (
	"encoding/json"
	"go-crud/models"
	"net/http"

	"gorm.io/gorm"
)

func CreateUser(writer http.ResponseWriter, reader *http.Request, db *gorm.DB) {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(reader.Body).Decode(&user)
	db.Create(&user)
}
