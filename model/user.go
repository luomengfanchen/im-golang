package model

import (
	"im/log"
	"time"
)

type User struct {
	Uid      int
	Token    string
	Email    string
	Name     string
	Password string
	CreateAt time.Time
}

// 验证用户
func AuthUser(valid string, passwd string) (User, error) {
	user := User{}

	// 从数据库中查找用户(一行记录)并填入结构体中
	sql := "SELECT uid, token, email, name, password, createAt FROM user_t WHERE email=$1 OR name = $1 AND password = $2"
	err := Db.QueryRow(sql, valid, passwd).
		Scan(&user.Uid, &user.Token, &user.Email, &user.Name, &user.Password, &user.CreateAt)

	return user, err
}

// 查询用户信息
func QueryUserRow(id int64) (User, error) {
	// 查询的用户信息
	var user User

	// 从数据库中查找用户
	err := Db.QueryRow("SELECT uid, token, email, name, password, createAt FROM user_t WHERE uid = $1", id).
		Scan(&user.Uid, &user.Token, &user.Email, &user.Name, &user.Password, &user.CreateAt)
	if err != nil {
		log.Warning(err.Error())
	}

	return user, err
}

// 查询多条用户信息
func QueryUser(id []int64) ([]User, error) {
	// 查询的用户信息
	var users []User
	var err error

	// 依次取出数据
	for _, v := range id {
		user := User{}
		// 从数据库中查找用户
		err = Db.QueryRow("SELECT uid, token, email, name, password, createAt FROM user_t WHERE uid = $1", v).
			Scan(&user.Uid, &user.Token, &user.Email, &user.Name, &user.Password, &user.CreateAt)
		if err != nil {
			log.Warning(err.Error())
		}

		// 保存数据到数组中
		users = append(users, user)
	}

	return users, err
}
