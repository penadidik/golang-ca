package main

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/handler"
	"user-service/repository"
	"user-service/usecase"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB("mongodb://localhost:27017")
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)

	router := mux.NewRouter()
	handler.NewUserHandler(router, userUC)

	log.Println("User service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
