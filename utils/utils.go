package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonData(w http.ResponseWriter, data interface{}, statusHeader int) {
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusHeader)
	w.Write(jsonData)
}
