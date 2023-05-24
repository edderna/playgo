package populators

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoPopulator interface {
	Populate(collection *mongo.Collection, file string)
}

type mongoPopulatorImpl struct {
}

func GetMongoPopulator() MongoPopulator {
	return new(mongoPopulatorImpl)
}

func (m *mongoPopulatorImpl) Populate(collection *mongo.Collection, file string) {
	byteData, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Can't populate %s with data file %s: %s", collection, file, err.Error())
		os.Exit(-1)
	}
	var data []interface{}
	unmarshalError := json.Unmarshal(byteData, &data)
	if unmarshalError != nil {
		fmt.Printf("Fail to unmarshal file %s: %s", file, err.Error())
		os.Exit(-1)
	}
	_, insertError := collection.InsertMany(context.Background(), data, options.InsertMany())
	if insertError != nil {
		fmt.Printf("Can't insert file %s into collection %s: %s", file, collection, err.Error())
		os.Exit(-1)
	}
}
