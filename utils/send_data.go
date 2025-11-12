package utils

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, product interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(product)

}
