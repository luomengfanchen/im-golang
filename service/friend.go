package service

import (
	"im/log"
	"im/model"
	"im/utils"
	"net/http"
	"strconv"
)

func Friend(w http.ResponseWriter, r *http.Request) {
	// 解析传递参数
	r.ParseForm()

	id := r.Form.Get("id")

	uid, _ := strconv.ParseInt(id, 10, 64)
	list, err := model.FriendList(uid)
	if err != nil {
		log.Warning(err.Error())
	}

	utils.ResponseJSON(w, http.StatusOK, utils.H{
		"data": list,
		"message": "success",
		"status": 200,
	})
}
