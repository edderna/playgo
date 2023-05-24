package repositories

import (
	"context"
	"cursoGo/internal/data_access"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const collectionName = "warehouses"

type WarehouseRepositoryMongoImpl struct {
	collection *mongo.Collection
}

func NewWarehouseRepositoryMongoImpl(m *mongo.Database) WarehouseRepositoryMongoImpl {
	return WarehouseRepositoryMongoImpl{
		collection: m.Collection(collectionName),
	}
}

func (w WarehouseRepositoryMongoImpl) FindByFFmCenterCode(ffmCenterCode string) (*data_access.Warehouse, error) {
	var result data_access.Warehouse
	var resultPointer = &result
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := w.collection.FindOne(ctx, bson.D{{
		Key:   "ffmCenterCode",
		Value: ffmCenterCode,
	}}).Decode(resultPointer)
	if err != nil {
		return nil, err
	}
	return resultPointer, nil
}
