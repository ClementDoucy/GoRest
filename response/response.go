package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data interface{}) {
	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
