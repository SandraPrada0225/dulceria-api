package entities

type Venta struct {
	ID            uint64
	MedioDePagoID uint64
	CarritoID     uint64
	CompradorID   uint64
	CreatedAt     uint64
}

func (Venta) TableName() string {
	return "ventas"
}
