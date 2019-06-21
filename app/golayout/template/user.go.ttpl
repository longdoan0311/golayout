package {{.ProjName}}

import (
	"time"
)

type User struct {
	ID 	string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	LastChangedAt time.Time `json:"lastChangedAt"`
}

type UserStore interface {
	GetAll() ([]User, error)
}

type UserInfoService interface {
	GetAll() ([]User, error)
}