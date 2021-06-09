package model

import (
	"im/log"
	"time")


type Chat struct {
	Cid      int
	SendId   int
	RecvId   int
	SendDate time.Time
	Msg      string
	UnRead   bool
}

// 取出聊天数据
func QueryChatAll(uid int64, fid int64) ([]Chat, error) {
	var list []Chat
	rows, err := Db.Query("SELECT cid, sendId, recvId, msg, unRead FROM chat_t WHERE sendId=$1 AND recvId=$2 OR sendID=$2 AND recvId=$1", uid, fid)
	if err != nil {
		log.Warning(err.Error())
	}

	for rows.Next() {
		chat := Chat{}
		err = rows.Scan(&chat.Cid, &chat.SendId, &chat.RecvId, &chat.SendDate, &chat.Msg, &chat.UnRead)
		list = append(list, chat)
	}

	return list, err
}