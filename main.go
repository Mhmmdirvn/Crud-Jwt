package main

import (
	"Combine-Gorm-Mux-Jwt/controllers"
	"Combine-Gorm-Mux-Jwt/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.HandlerLogin)

	handler := middlewares.CorsMiddleware().Handler(middlewares.MiddlewareJWTAuthorization(router))
	
	router.HandleFunc("/users", controllers.ReadPersons).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.ReadPersonById).Methods("GET")
	router.HandleFunc("/create", controllers.AddPerson).Methods("POST")
	router.HandleFunc("/edit/{id}", controllers.EditPerson).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePerson).Methods("DELETE")

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", handler)

}