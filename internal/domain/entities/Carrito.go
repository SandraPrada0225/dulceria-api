package entities

type Carrito struct {
	ID              uint64
	Subtotal        float64
	Descuento       float64
	Envio           float64
	PrecioTotal     float64
	EstadoCarritoID uint64
}
