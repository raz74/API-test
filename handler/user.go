package handler

import (
	"api-test/models"
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const BASE_URL = "http://127.0.0.1:3000"

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

	_, _ = p.create(req, file)
	//return c.JSON(http.StatusOK, &newUser)
	return c.Redirect(http.StatusOK, "http://127.0.0.1:9000/")
}

func (p *PostgresRepo) create(req models.UserReq, file *multipart.FileHeader) (*models.User, error) {

	newUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		PhotoName: file.Filename,
		CreatedAt: time.Now(),
		Status:    "active",
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
		return echo.ErrBadRequest
	}
	//imageUrl := BASE_URL + user.PhotoName

	return c.JSON(http.StatusOK, &user)
}

func (p *PostgresRepo) GetUsers(c echo.Context) error {

	var users []models.User
	err := p.DB.Where("status = ?", "active").Find(&users).Error
	if err != nil {
		return echo.ErrBadRequest
	}
	type OutputStruct struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Status   string `json:"status"`
		ImageUrl string `json:"image_url"`
	}
	var outputData []OutputStruct
	for _, user := range users {
		if len(user.Name) > 0 {
			outputData = append(outputData, OutputStruct{
				Id:       user.Id,
				Name:     user.Name,
				Email:    user.Email,
				Status:   user.Status,
				ImageUrl: BASE_URL + "/static/" + user.PhotoName,
			})
		}
	}
	type userOutput struct {
		Data []OutputStruct `json:"data"`
	}
	err = c.JSON(http.StatusOK, &userOutput{Data: outputData})
	if err != nil {
		fmt.Println(err)
	}
	return err
}
