package users_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateUser(writer http.ResponseWriter, reader *http.Request, db *gorm.DB) {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(reader.Body).Decode(&user)
	fmt.Println(reader.Body)
	fmt.Println(user)
	db.Create(&user)
}

func GetUser(writer http.ResponseWriter, reader *http.Request, db *gorm.DB) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(reader)
	userId := params["id"]

	var userFound models.User

	result := db.First(&userFound, userId)
	if result.Error != nil { //Error handling
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(writer, "User not found", http.StatusNotFound)
		} else {
			http.Error(writer, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}
	writer.WriteHeader(http.StatusOK)

	userJSON, err := json.Marshal(userFound)
	if err != nil { // Error handling
		http.Error(writer, "Failed to encode user", http.StatusInternalServerError)
		return
	}
	writer.Write(userJSON)
}

func UpdateUser(writer http.ResponseWriter, reader *http.Request, db *gorm.DB) {
	params := mux.Vars(reader)
	userId := params["id"]
	var userFound models.User
	result := db.First(&userFound, userId)
	if result.Error != nil { //Error handling
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(writer, "User for update not found", http.StatusNotFound)
		} else {
			http.Error(writer, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}
	var userUpdate models.User
	_ = json.NewDecoder(reader.Body).Decode(&userUpdate)
	fmt.Println(userFound)
	userFound = userUpdate
	fmt.Println(userUpdate)
}
