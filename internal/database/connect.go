package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

func ConnectToDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Error("dont connect to mongoDB")
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Error("dont connect to mongoDB")
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error("database dont ping")
		log.Fatal(err)
	}

	log.Info("Connection to database success")
	return client
}

func ConnectClose(client *mongo.Client) {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Error("database dont close connection")
		log.Fatal(err)
	}
	log.Info("Connection to database closed")
}
