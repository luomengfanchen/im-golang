package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

func ResponseJSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)

	ret, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(ret)
}