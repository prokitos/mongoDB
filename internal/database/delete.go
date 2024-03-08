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

	bdoc := jsonToInterf(DBmodel)

	deleteResult, err := collection.DeleteMany(context.TODO(), bdoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
