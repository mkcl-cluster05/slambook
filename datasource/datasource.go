package datasource

import (
	"fmt"
	"slambook/utils/config"
	"slambook/utils/dbhelper"

	"go.mongodb.org/mongo-driver/mongo"
)

type DataSource struct {
	MongoDB *mongo.Client
}

func InitDS() (*DataSource, error) {

	mongoDB, err := dbhelper.GetMongoConnection(config.DatabaseConfig.Mongo.DSN)
	if err != nil {
		return nil, fmt.Errorf("error opening mongodb : %w", err)
	}

	return &DataSource{
		MongoDB: mongoDB,
	}, nil

}
