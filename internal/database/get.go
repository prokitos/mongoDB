package database

import (
	"context"
	"fmt"
	"log"
	"module/internal/models"
)

func SearchToken(dbModel models.TokenDB) models.TokenDB {

	client := ConnectToDB()
	defer ConnectClose(client)

	collection := client.Database("mongo").Collection("tokenJWT")

	var result models.TokenDB

	bdoc := jsonToInterf(dbModel)

	err := collection.FindOne(context.TODO(), bdoc).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result
}
