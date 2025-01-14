package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var products []Product

func main() {

	err := InitSliceProduct()

	if err != nil {
		errors.New("error initialized sliceProduct")
	}

	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.Id, product.Name, product.Price)
	}

}

func InitSliceProduct() error {

	file, err := os.Open("products.json")
	if err != nil {
		log.Fatalf("Error open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error al leer el archivo: %v", err)
	}
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		log.Fatalf("Error al decodificar JSON: %v", err)
	}

	return nil
}
