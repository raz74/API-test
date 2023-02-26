package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int            `json:"-"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Photo     []byte         `form:"photo" json:"-"`
	Status    string         `json:"status"`
	PhotoName string         `json:"photoName"`
}

type UserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo []byte `form:"photo"`
}
