package service

import (
	"im/log"
	"im/model"
	"im/utils"
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
