package app

import (
	"encoding/json"
	"module/internal/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func useToken(w http.ResponseWriter, r *http.Request) {

	// получение токена из запроса, и тестирование его использования
	var bearerToken string = ""
	bearerToken = r.Header.Get("Authorization")
	response := service.TokenAccessValidate(bearerToken)

	json.NewEncoder(w).Encode(response)
	log.Info("user use token to access")
}
