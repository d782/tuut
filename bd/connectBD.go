package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = connectBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://diegocano:qwer1234@cluster0.n9ajk.mongodb.net/dibs?retryWrites=true&w=majority")

func connectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connection Successfull!")
	return client
}

func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
