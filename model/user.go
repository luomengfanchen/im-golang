package model

import "time"

type User struct {
	Uid      string
	Uuid     string
	Email    string
	Name     string
	Password string
	CreateAt time.Time
}
