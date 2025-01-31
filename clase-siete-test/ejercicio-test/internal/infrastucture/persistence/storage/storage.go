package storage

import "github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain"

type Storage interface {
	ReadAll() ([]domain.Product, error)
	WriteAll([]domain.Product) error
}
