package storage

import (
	"encoding/json"
	"io"
	"os"

	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain"
)

type StorageJSON struct {
	filePath string
}

func NewStorageJSON(filePath string) StorageJSON {
	return StorageJSON{
		filePath: filePath,
	}
}

func (strorage *StorageJSON) ReadAll() ([]domain.Product, error) {
	file, err := os.Open(strorage.filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []domain.Product

	reader, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(reader, &products)
	if err != nil {
		return nil, err
	}
	return products, nil

}

func (storageJson *StorageJSON) WriteAll(products []domain.Product) error {
	jsonProducts, err := json.MarshalIndent(products, "", "  ")

	if err != nil {
		return err
	}

	file, err := os.Create("../../products.json")

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonProducts)
	if err != nil {
		return err
	}

	return nil

}
