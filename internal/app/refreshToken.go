package app

import (
	"encoding/json"
	"module/internal/service"
	"net/http"
)

func refreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// получение рефреш токена из запроса, и получение нового аксес токена
	refreshToken := r.Header.Get("Refresher")
	newAccessToken := service.RenewToken(refreshToken)

	// вывод нового токена пользователю
	json.NewEncoder(w).Encode(newAccessToken)
}
