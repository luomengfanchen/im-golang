package model

import "im/log"

type Friend struct {
	Uid int
	Fid int
}

// 查询该用户的所有好友
func QueryFriendAll(id int64) ([]Friend, error) {
	// 返回的好友列表数组
	var friends []Friend

	// 查询好友
	rows, err := Db.Query("SELECT * FROM friend_t WHERE uid = $1", id)
	if err != nil {
		log.Warning(err.Error())
	}

	// 将数据填入
	for rows.Next() {
		friend := Friend{}
		err = rows.Scan(&friend.Uid, &friend.Fid)
		friends = append(friends, friend)
	}

	return friends, err
}