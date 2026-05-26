package purchasecarrito

import "dulceria-api/internal/domain/dto/contracts/purchasecarrito"

//intruccion de negocio//aplana la estructura, convierte
//datos de html en datos de negocio

type PurchasecarritoCommand struct {
	CarritoID     uint64
	CompradorID   uint64
	MedioDePagoID uint64
}

func NewPurchaseCarritoCommandFromRequest(request purchasecarrito.Request) PurchasecarritoCommand {
	return PurchasecarritoCommand{
		CompradorID:   request.Body.CompradorID,
		CarritoID:     request.URLParams.CarritoID,
		MedioDePagoID: request.Body.MedioDePagoID,
	}
}
