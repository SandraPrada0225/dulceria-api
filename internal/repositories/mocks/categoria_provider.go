package mocks

import (
	"dulceria-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCategoriaProvider struct {
	mock.Mock
}

func (mock *MockCategoriaProvider) GetAll() ([]entities.Categoria, error) {
	args := mock.Called()
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]entities.Categoria), err
	}
	return []entities.Categoria{}, err
}

func (m *MockCategoriaProvider) GetCategoriasByDulceID(dulceID uint64) ([]entities.Categoria, error) {
	resposeArgs := m.Called(dulceID)
	response := resposeArgs.Get(0)
	err := resposeArgs.Error(1)
	return response.([]entities.Categoria), err
}
