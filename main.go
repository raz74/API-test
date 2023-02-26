package main

import (
	"api-test/handler"
	"api-test/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/labstack/gommon/log"
)

func main() {
	db := repository.Initialize()
	h := handler.NewPostgresRepo(db)

	setUpRouts(h)
}

func setUpRouts(h *handler.PostgresRepo) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/static", "static")
	e.POST("/user", h.UploadFile)
	e.GET("/users/", h.GetUsers)
	e.GET("/user/:id", h.GetUser)
	e.Logger.Fatal(e.Start(":3000"))
}
