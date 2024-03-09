package database

import (
	"context"
	"fmt"
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func DeleteToken(DBmodel models.TokenDB) {

	client := ConnectToDB()
	defer ConnectClose(client)

	collection := client.Database("mongo").Collection("tokenJWT")

	bdoc := jsonToInterf(DBmodel)

	deleteResult, err := collection.DeleteMany(context.TODO(), bdoc)
	if err != nil {
		log.Error("error convert json to interface")
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
