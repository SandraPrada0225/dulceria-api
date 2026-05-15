package entities

type Usuario struct {
	ID              uint64
	Nombre          string
	Apellido        string
	Password        string
	correo          string
	CarritoActualID uint64
}
