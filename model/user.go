package model

import "time"

type User struct {
	Uid      int
	Token    string
	Email    string
	Name     string
	Avator   string
	Password string
	CreateAt time.Time
}

// 查找用户
func Find(id int) (User, error) {
	user := User{}

	// 从数据库中查找用户(一行记录)并填入结构体中
	err := Db.QueryRow("SELECT uid, token, email, name, avator, password, createAt FROM user_t WHERE uid = $1", id).
		Scan(&user.Uid, &user.Token, &user.Email, &user.Name, &user.Avator, &user.Password, &user.CreateAt)

	return user, err
}

// 验证用户是否存在
func AuthUser(email string, password string) (bool) {
	ok := Db.QueryRow("SELECT * FROM user_t WHERE email=$1 OR name=$1 AND password=$2", email, password).Err()

	if ok != nil {
		return false
	} else {
		return true
	}
}
