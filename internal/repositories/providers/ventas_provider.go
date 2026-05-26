package providers

import (
	"dulceria-api/internal/domain/dto/responses"
	"dulceria-api/internal/domain/entities"
)

type VentasProvider interface {
	Create(venta *entities.Venta) error
	GetListByUserId(userID uint64) (responses.GetPurchaseList, error)
}
