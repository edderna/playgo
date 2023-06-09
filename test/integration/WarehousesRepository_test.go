package integration

import (
	"cursoGo/internal/warehouses"
	"cursoGo/test/test_utils/populators"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
	"testing"
)

func initContext() warehouses.WarehouseRepository {
	populators.GetMongoPopulator().Populate(context.mongoClient.Database("test").
		Collection("warehouses"), "../../resources/warehouse.json")
	return warehouses.NewWarehouseRepository(context.mongoClient.Database("test"))
}

func TestFindByFfmCenterCode_whenExistsFfmCenter_mustReturnFfmCenter(t *testing.T) {
	repository := initContext()
	result, err := repository.FindByFFmCenterCode("ZARAESPAR")
	if result == nil {
		t.Fatalf("%s", err.Error())
	} else {
		assert.DeepEqual(t, *result, warehouses.Warehouse{Code: "437", FfmCenterCode: "ZARAESPAR", BrandId: "1", Name: "Zara – Parla", City: "Parla",
			AddressLine1: "Calle Nantes", ZipCode: "28983", Province: "Madrid", Country: "España"}, cmpopts.IgnoreFields(*result, "Id"))
	}
}
