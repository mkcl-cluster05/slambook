package dbhelper

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *mongo.Client
var sessionError error
var once sync.Once

func GetMongoConnection(mongoDSN string) (*mongo.Client, error) {

	once.Do(func() {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(mongoDSN)

		clientOptions.SetMaxConnIdleTime(100)

		clientOptions.SetMaxPoolSize(1000)

		clientOptions.SetMaxConnIdleTime(4 * time.Hour)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			sessionError = err
			return
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			sessionError = err
			return
		}
		instance = client
	})

	return instance, sessionError

}
