package entities

type CarritoDulce struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	CarritoID uint64  `json:"carrito_id"`
	DulceID   uint64  `json:"dulce_id"`
	Unidades  int     `json:"unidades"`
	Subtotal  float64 `json:"subtotal"`
}
