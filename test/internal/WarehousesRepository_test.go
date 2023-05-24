package internal

import (
	"cursoGo/internal/data_access/repositories"
	repositoriesImpl "cursoGo/internal/data_access/repositories/impl"
	"cursoGo/test/test_utils/populators"
	"fmt"
	"testing"
)

func initContext() repositories.WarehouseRepository {
	populators.GetMongoPopulator().Populate(context.mongoClient.Database("test").
		Collection("warehouses"), "../../resources/warehouse.json")
	return repositoriesImpl.NewWarehouseRepositoryMongoImpl(context.mongoClient.Database("test"))
}

func TestFindByFfmCenterCode_whenExistsFfmCenter_mustReturnFfmCenter(t *testing.T) {
	repository := initContext()
	result, err := repository.FindByFFmCenterCode("ZARAESPAR")
	if result == nil {
		t.Fatalf("%s", err.Error())
	} else {
		fmt.Println(result)
	}
}
