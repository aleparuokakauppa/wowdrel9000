package db

import (
    "context"
    "log"
    "sync"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    clientInstanceError error
    mongoOnce sync.Once
)

func GetMongoClient(ctx context.Context) (*mongo.Client, error) {
    mongoOnce.Do(func() {
        clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

        client, err := mongo.Connect(ctx, clientOptions)
        if err != nil {
            clientInstanceError = err
            return
        }

        err = client.Ping(context.TODO(), nil)
        if err != nil {
            clientInstanceError = err
            return
        }

        clientInstance = client
    })
    return clientInstance, clientInstanceError
}

func CloseMongoClient(ctx context.Context) {
    if clientInstance != nil {
        err := clientInstance.Disconnect(ctx)
        if err != nil {
            log.Fatal(err)
        }
    }
}
