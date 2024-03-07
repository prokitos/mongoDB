package database

import (
	"context"
	"fmt"
	"log"
	"module/internal/models"
)

func DeleteToken(DBmodel models.TokenDB) {

	client := ConnectToDB()
	defer ConnectClose(client)

	collection := client.Database("mongo").Collection("tokenJWT")

	deleteResult, err := collection.DeleteOne(context.TODO(), DBmodel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
