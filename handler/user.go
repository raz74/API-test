package handler

import (
	"api-test/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type PostgresRepo struct {
	DB *gorm.DB
}

func NewPostgresRepo(DB *gorm.DB) *PostgresRepo {
	return &PostgresRepo{DB: DB}
}

func (p *PostgresRepo) UploadFile(c echo.Context) error {
	var req models.UserReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	// Destination
	dst, err := os.Create("static/" + file.Filename)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	newUser, err := p.create(req, file)
	return c.JSON(http.StatusOK, &newUser)
}

func (p *PostgresRepo) create(req models.UserReq, file *multipart.FileHeader) (*models.User, error) {

	newUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		PhotoName: file.Filename,
	}

	err := p.DB.Create(newUser).Error
	if err != nil {
		return nil, echo.ErrBadRequest
	}
	return newUser, err
}

func (p *PostgresRepo) GetUser(c echo.Context) error {

	id := c.Param("id")

	var user *models.User
	err := p.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &user)
}
