package data_access

type Warehouse struct {
	Id            string `json:"id" bson:"_id"`
	Code          string `json:"code"`
	FfmCenterCode string `json:"ffmCenterCode"`
	BrandId       string `json:"brandId"`
	Name          string `json:"name"`
	AddressLine1  string `json:"addressLine1"`
	ZipCode       string `json:"zipCode"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
}

type WarehouseRepository interface {
	FindByFFmCenterCode(ffmCenterCode string) (*Warehouse, error)
}
