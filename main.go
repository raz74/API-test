package main

import (
	"api-test/handler"
	"api-test/repository"
	"github.com/labstack/echo"
)

func main() {
	db := repository.Initialize()
	h := handler.NewPostgresRepo(db)

	setUpRouts(h)
}

func setUpRouts(h *handler.PostgresRepo) {
	e := echo.New()
	e.Static("/static", "static")
	e.POST("/photo", h.UploadFile)
	e.GET("/photo/:id", h.GetUser)
	e.Logger.Fatal(e.Start(":3000"))
}
