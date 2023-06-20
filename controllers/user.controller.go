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

func UpdateUser(writer http.ResponseWriter, reader *http.Request, db *gorm.DB) { // need testing
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(reader)
	userID := params["id"]

	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(writer, "User not found", http.StatusNotFound)
		} else {
			http.Error(writer, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(reader.Body).Decode(&updatedUser); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email

	if err := db.Save(&existingUser).Error; err != nil {
		http.Error(writer, "Failed to update user", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(existingUser)
}

func DeleteUser(writer http.ResponseWriter, reader *http.Request, db *gorm.DB) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(reader)
	userID := params["id"]

	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(writer, "User not found", http.StatusNotFound)
		} else {
			http.Error(writer, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	existingUser.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := db.Save(&existingUser).Error; err != nil {
		http.Error(writer, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(existingUser)
}
