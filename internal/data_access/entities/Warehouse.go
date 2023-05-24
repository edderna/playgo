package entities

type Warehouse struct {
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
