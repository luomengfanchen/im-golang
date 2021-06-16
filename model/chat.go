package model

import (
	"im/log"
	"time"
)

type Chat struct {
	Cid      int
	SendId   int
	RecvId   int
	SendDate time.Time
	Msg      string
	UnRead   bool
}

// 取出聊天数据
func QueryChat(uid int64, fid int64) ([]Chat, error) {
	var list []Chat
	rows, err := Db.Query("SELECT cid, sendId, recvId, sendDate, msg, unRead FROM chat_t WHERE sendId=$1 AND recvId=$2 OR sendID=$2 AND recvId=$1", uid, fid)
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

// 插入聊天记录
func InsertChatRow(sendId int, recvId int, msg string) error {
	stmt, err := Db.Prepare("INSERT INTO chat_t (sendId, recvId, sendDate, msg, unRead) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	err = stmt.QueryRow(sendId, recvId,nowTime, msg, true).Err()
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}
