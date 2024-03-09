package app

import (
	"encoding/json"
	"module/internal/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func refreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// получение рефреш токена из запроса
	var refreshToken string = ""
	refreshToken = r.Header.Get("Refresher")
	// получение аксес токена из запроса
	var bearerToken string = ""
	bearerToken = r.Header.Get("Authorization")

	// получение нового аксес и рефреш токена
	newAccessToken := service.RenewToken(refreshToken, bearerToken)

	// вывод нового токена пользователю
	json.NewEncoder(w).Encode(newAccessToken)
	log.Info("user refresh tokens")
}
