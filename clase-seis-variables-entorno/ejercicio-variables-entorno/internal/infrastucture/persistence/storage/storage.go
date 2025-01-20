package storage

import "github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/domain"

type Storage interface {
	ReadAll() ([]domain.Product, error)
	WriteAll([]domain.Product) error
}
