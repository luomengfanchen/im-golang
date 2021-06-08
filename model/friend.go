package model

import "im/log"

type Friend struct {
	Uid int
	Fid int
}

func FriendList(id int64) ([]Friend, error){
	var list []Friend
	rows, err := Db.Query("SELECT * FROM friend_t WHERE uid = $1", id)
	if err != nil {
		log.Warning(err.Error())
	}

	for rows.Next() {
		friend := Friend{}
		err = rows.Scan(&friend.Uid, &friend.Fid)
		list = append(list, friend)
	}

	return list, err
}