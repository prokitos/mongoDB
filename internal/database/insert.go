package database

import (
	"context"
	"fmt"
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func InsertToken(DBmodel models.TokenDB) {

	client := ConnectToDB()
	defer ConnectClose(client)

	collection := client.Database("mongo").Collection("tokenJWT")
	bdoc := jsonToInterf(DBmodel)

	insertManyResult, err := collection.InsertOne(context.TODO(), bdoc)
	if err != nil {
		log.Error("error execute command")
		log.Fatal(err)
	}

	fmt.Println(insertManyResult.InsertedID)

}
