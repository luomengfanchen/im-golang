package service

import (
	"im/log"
	"im/model"
	"im/utils"
	"net/http"
	"strconv"
)

type friend struct {
	Uid  int
	Name string
}

func Friend(w http.ResponseWriter, r *http.Request) {
	// 解析传递参数
	r.ParseForm()

	id := r.Form.Get("id")

	// 查找当前用户好友
	uid, _ := strconv.ParseInt(id, 10, 64)
	fids, err := model.QueryFriend(uid)
	if err != nil {
		log.Warning(err.Error())
	}

	// 添加当前用户好友
	users, _ := model.QueryUser(fids)
	var list []friend
	for _, v := range users {
		u := friend{
			Uid:  v.Uid,
			Name: v.Name,
		}
		list = append(list, u)
	}

	utils.ResponseJSON(w, http.StatusOK, utils.H{
		"data":    list,
		"message": "success",
		"status":  200,
	})
}
