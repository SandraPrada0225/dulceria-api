package providers

import "dulceria-api/internal/domain/entities"

type MarcasProvider interface {
	GetAll() (marca []entities.Marca, err error)
}
