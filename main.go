package main

import (
	"log"

	"github.com/buniekbua/gousers/db"
	"github.com/buniekbua/gousers/handlers"
	"github.com/buniekbua/gousers/repositories"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repositories.NewUserRepository(db)

	userHandler := handlers.NewUserHandler(userRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", userHandler.GetAllUsers)
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
