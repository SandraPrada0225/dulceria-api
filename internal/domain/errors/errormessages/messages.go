package errormessages

import "fmt"

type (
	ErrorMessage string
	Parameters   map[string]interface{}
)

const (
	CarrritoDulceNotFound          ErrorMessage = "no se encontró un detalle carrito_dulce con este código"
	DulceNotFound                  ErrorMessage = "No se encontró un dulce con ese código"
	UsuarioNotFound                ErrorMessage = "No se encontró usuario con ese id"
	CarritoNotBelonging            ErrorMessage = "El carrito pertenece a otro usuario"
	InternalServerError            ErrorMessage = "Ha ocurrido un error inesperado"
	IdMustBeAPositiveNumber        ErrorMessage = "El ID debe ser un número positivo"
	UnitLimitExceeded              ErrorMessage = "Las unidades requeridas esxeden las disponibles"
	CarritoNotFound                ErrorMessage = "No se encontró un carrito con ese id"
	InvalidTypeError               ErrorMessage = "El tipo de dato es invalido"
	CarritoHasAlreadyBeenPurchased ErrorMessage = "El carrito ya ha sido comprado"
)

func (e ErrorMessage) String() string {
	return string(e)
}

func (e ErrorMessage) GetMessageWithParams(params Parameters) string {
	msg := e.String()

	for key, value := range params {
		msg = fmt.Sprintf("%s. %s: %s", msg, key, value)
	}
	return msg
}
