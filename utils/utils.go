package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonData(w http.ResponseWriter, data interface{}, statusHeader int) {
	jsonData, _ := json.Marshal(data)
	// jsonData by unmarshal : [123 34 105 100 34 58 49 52 44 34 98 111 111 107 110 97 109 101 34 58
	// 34 97 115 100 97 115 100 97 115 100 34 44 34 112 117 98 108 105 99 97
	// 116 105 111 110 115 34 58 34 115 97 100 107 108 106 97 115 108 107 100 106 97 115 108 107 106 100 97 108 107 115 106 100 108 107 97 115 106
	// 100 97 115 34 125]
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusHeader)
	w.Write(jsonData)
}
