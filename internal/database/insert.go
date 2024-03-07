package database

import (
	"context"
	"fmt"
	"log"
	"module/internal/models"
)

func InsertToken(DBmodel models.TokenDB) {

	client := ConnectToDB()
	defer ConnectClose(client)

	collection := client.Database("mongo").Collection("tokenJWT")

	insertManyResult, err := collection.InsertOne(context.TODO(), DBmodel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertManyResult.InsertedID)

}
