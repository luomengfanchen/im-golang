package service

import (
	"encoding/json"
	"im/log"
	"im/model"
	"im/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ChatList(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	r.ParseForm()

	uid, err := strconv.ParseInt(r.Form.Get("uid"), 10, 64)
	if err != nil {
		log.Warning(err.Error(), uid)
	}
	fid, err := strconv.ParseInt(r.Form.Get("fid"), 10, 64)
	if err != nil {
		log.Warning(err.Error(), fid)
	}

	chatList, err := model.QueryChat(uid, fid)
	if err != nil {
		log.Warning(err.Error())
	}

	utils.ResponseJSON(w, http.StatusOK, utils.H{
		"data":    chatList,
		"message": "success",
		"status":  200,
	})
}

// 发送消息的json结构体
type msg struct {
	RecvId int
	SendId int
	Msg    string
}

// 发送消息
func Chat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Info(err.Error())
		return
	}
	recvmsg := msg{}
	json.Unmarshal(body, &recvmsg)

	err = model.InsertChatRow(recvmsg.SendId, recvmsg.RecvId, recvmsg.Msg)
	if err != nil {
		log.Warning(err.Error())
	}
}
