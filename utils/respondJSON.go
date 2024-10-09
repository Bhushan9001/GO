package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter ,stats int , payload interface{}){
	response,_ := json.Marshal(payload);

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stats);
	w.Write(response);
}