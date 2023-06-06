package router

import (
	users_controller "go-crud/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func User_router(db *gorm.DB) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users/create", func(writer http.ResponseWriter, reader *http.Request) {
		users_controller.CreateUser(writer, reader, db)
	}).Methods("POST")
	return r
}
