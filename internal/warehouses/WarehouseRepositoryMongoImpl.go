package warehouses

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type WarehouseRepositoryMongoImpl struct {
	Collection *mongo.Collection
}

func (w WarehouseRepositoryMongoImpl) FindByFFmCenterCode(ffmCenterCode string) (*Warehouse, error) {
	var result Warehouse
	var resultPointer = &result
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := w.Collection.FindOne(ctx, bson.D{{
		Key:   "ffmCenterCode",
		Value: ffmCenterCode,
	}}).Decode(resultPointer)
	if err != nil {
		return nil, err
	}
	return resultPointer, nil
}
