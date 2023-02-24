package main

import (
	"api-test/handler"
	"api-test/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := repository.Initialize()
	h := handler.NewPostgresRepo(db)

	setUpRouts(h)
}

func setUpRouts(h *handler.PostgresRepo) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/static", "static")
	e.POST("/photo", h.UploadFile)
	e.GET("/photos/", h.GetUsers)
	e.GET("/photo/:id", h.GetUser)
	e.Logger.Fatal(e.Start(":3000"))
}
