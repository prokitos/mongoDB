package database

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// перевод json в interface (чтобы пустые поля из json не передавались в запрос)
func jsonToInterf(temp interface{}) interface{} {

	out, _ := json.MarshalIndent(temp, "", "  ")
	fmt.Print(string(out))

	var bdoc interface{}
	err := bson.UnmarshalExtJSON([]byte(out), true, &bdoc)
	if err != nil {
		panic(err)
	}

	return bdoc
}
