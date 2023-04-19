package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yizumi/go-test/controllers"
	"github.com/yizumi/go-test/models"
	"github.com/yizumi/go-test/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	r := mux.NewRouter()
	userController.RegisterRoutes(r)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
