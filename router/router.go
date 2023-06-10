package router

import (
	users_controller "go-crud/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func User_router(db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	//CREATE USER
	r.HandleFunc("/users/create", func(writer http.ResponseWriter, reader *http.Request) {
		users_controller.CreateUser(writer, reader, db)
	}).Methods("POST")

	//GET USER
	r.HandleFunc("/users/{id}", func(writer http.ResponseWriter, reader *http.Request) {
		users_controller.GetUser(writer, reader, db)
	}).Methods("GET")

	/*//UPDATE USER
	r.HandleFunc("/users/update", func(writer http.ResponseWriter, reader *http.Request) {
		users_controller.UpdateUser(writer, reader, db)
	}).Methods("PATCH")
	*/

	return r
}
