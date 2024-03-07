package app

import (
	"encoding/json"
	"module/internal/service"
	"net/http"
)

func useToken(w http.ResponseWriter, r *http.Request) {

	// получение токена из запроса, и тестирование его использования
	bearerToken := r.Header.Get("Authorization")
	response := service.TokenAccessValidate(bearerToken)

	json.NewEncoder(w).Encode(response)
}
