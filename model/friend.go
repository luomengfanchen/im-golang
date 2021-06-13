package model

import "im/log"

type Friend struct {
	Uid int
	Fid int
}

// 查询该用户的所有好友
func QueryFriend(id int64) ([]int64, error) {
	// 返回的好友列表数组
	var fids []int64

	// 查询好友
	rows, err := Db.Query("SELECT fid FROM friend_t WHERE uid = $1", id)
	if err != nil {
		log.Warning(err.Error())
	}

	// 将数据填入
	for rows.Next() {
		var fid int64
		err = rows.Scan(&fid)
		fids = append(fids, fid)
	}

	return fids, err
}