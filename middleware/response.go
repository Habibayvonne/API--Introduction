package middleware

import (
	"encoding/json"
	"net/http"
)

// JSONResponse takes any response from a controller and turns it into a json
// with the structure
/*
{
	"data": data
}
*/

func JSONResponse(w http.ResponseWriter, code int, data interface{}) {
	responseJSON, _ := json.Marshal(data)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(responseJSON)
}
