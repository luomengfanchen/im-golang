package model

import "time"

type User struct {
	Uid      int
	Token    string
	Email    string
	Name     string
	Password string
	CreateAt time.Time
}

// 验证用户
func Auth(valid string, passwd string) (User, error) {
	user := User{}

	// 从数据库中查找用户(一行记录)并填入结构体中
	sql := "SELECT uid, token, email, name, password, createAt FROM user_t WHERE email=$1 OR name = $1 AND password = $2"
	err := Db.QueryRow(sql, valid, passwd).
		Scan(&user.Uid, &user.Token, &user.Email, &user.Name, &user.Password, &user.CreateAt)

	return user, err
}
