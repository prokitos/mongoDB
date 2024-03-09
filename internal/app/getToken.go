package app

import (
	"encoding/json"
	"module/internal/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func getToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// получение guid по запросу
	var GUID string = ""
	GUID = r.FormValue("GUID")

	// отправка guid в токен
	tokens := service.TokenGetPair(GUID)

	// вывод токенов пользователю
	json.NewEncoder(w).Encode(tokens)
	log.Info("user get new tokens")
}
