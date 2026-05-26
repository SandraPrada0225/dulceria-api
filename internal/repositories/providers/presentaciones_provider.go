package providers

import "dulceria-api/internal/domain/entities"

type PresentacionesProvider interface {
	GetAll() (presentacion []entities.Presentacion, err error)
}
