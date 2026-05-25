package business

type UnitLimitExceeded struct {
	mensaje string
}

func (e UnitLimitExceeded) Error() string {
	return e.mensaje
}

func NewUnitLimitExceeded(mensaje string) error {
	return UnitLimitExceeded{
		mensaje: mensaje,
	}
}
