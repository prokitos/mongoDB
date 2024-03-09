package database

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// перевод json в interface (чтобы пустые поля из json не передавались в запрос)
func jsonToInterf(temp interface{}) interface{} {

	out, _ := json.MarshalIndent(temp, "", "  ")

	var bdoc interface{}
	err := bson.UnmarshalExtJSON([]byte(out), true, &bdoc)
	if err != nil {
		log.Error("error unmarshal json")
		panic(err)
	}

	return bdoc
}
