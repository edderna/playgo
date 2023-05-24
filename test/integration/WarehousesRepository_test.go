package integration

import (
	"cursoGo/internal/data_access"
	"cursoGo/internal/data_access/impl"
	"cursoGo/test/test_utils/populators"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
	"testing"
)

func initContext() data_access.WarehouseRepository {
	populators.GetMongoPopulator().Populate(context.mongoClient.Database("test").
		Collection("warehouses"), "../../resources/warehouse.json")
	return repositories.NewWarehouseRepositoryMongoImpl(context.mongoClient.Database("test"))
}

func TestFindByFfmCenterCode_whenExistsFfmCenter_mustReturnFfmCenter(t *testing.T) {
	repository := initContext()
	result, err := repository.FindByFFmCenterCode("ZARAESPAR")
	if result == nil {
		t.Fatalf("%s", err.Error())
	} else {
		assert.DeepEqual(t, *result, data_access.Warehouse{Code: "437", FfmCenterCode: "ZARAESPAR", BrandId: "1", Name: "Zara – Parla", City: "Parla",
			AddressLine1: "Calle Nantes", ZipCode: "28983", Province: "Madrid", Country: "España"}, cmpopts.IgnoreFields(*result, "Id"))
	}
}
