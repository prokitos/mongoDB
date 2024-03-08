package app

import (
	"encoding/json"
	"module/internal/service"
	"net/http"
)

func getToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// получение guid по запросу
	GUID := r.FormValue("GUID")

	// отправка guid для получения токенов
	tokens := service.TokenGetPair(GUID)

	// вывод токенов пользователю
	json.NewEncoder(w).Encode(tokens)
}
