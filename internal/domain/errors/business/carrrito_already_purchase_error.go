package business

// un carrito ya comprado no puede volver a comprarse
type CarritoAlreadyPurchaseError struct {
	mensaje string
}

func (e CarritoAlreadyPurchaseError) Error() string {
	return e.mensaje
}

func NewCarritoAlreadyPurchaseError(mensaje string) error {
	return CarritoAlreadyPurchaseError{
		mensaje: mensaje,
	}
}
