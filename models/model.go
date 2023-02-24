package models

import "time"

type User struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Photo     []byte    `form:"photo" json:"-"`
	Status    string    `json:"status"`
	PhotoName string    `json:"photoName"`
}

type UserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo []byte `form:"photo"`
}
