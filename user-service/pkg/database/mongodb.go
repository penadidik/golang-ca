package database

import (
    "context"
    "log"
    "time"

    "https://github.com/penadidik/golang-ca/tree/7f1eacc2fa3886a87280f5417c914085ced10d88/user-service/config"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Database {
    client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoURI()))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    return client.Database("userdb")
}
