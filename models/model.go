package models

type User struct {
	Id        int
	Name      string `json:"name"`
	Email     string `json:"email"`
	Photo     []byte `form:"photo"`
	PhotoName string `json:"photoName"`
}

type UserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo []byte `form:"photo"`
}
