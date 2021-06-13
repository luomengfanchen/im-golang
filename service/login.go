package service

import (
	"im/model"
	"im/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	r.ParseForm()

	// 获取参数
	name := r.Form.Get("username")
	passwd := r.Form.Get("password")

	user, err := model.AuthUser(name, passwd)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.H{
			"message": "failure",
			"status":  "400",
		})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, utils.H{
		"data": utils.H{
			"uid":   user.Uid,
			"name":  user.Name,
			"email": user.Email,
			"token": user.Token,
		},
		"message": "success",
		"status":  200,
	})
}
