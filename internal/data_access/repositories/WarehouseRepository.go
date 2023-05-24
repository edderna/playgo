package repositories

import "cursoGo/internal/data_access/entities"

type WarehouseRepository interface {
	FindByFFmCenterCode(ffmCenterCode string) (*entities.Warehouse, error)
}
