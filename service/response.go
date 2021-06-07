package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	ret, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(ret)
}