package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

	router := chi.NewRouter()

	router.Get("/ping", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("pong"))
		response.WriteHeader(http.StatusOK)
	})

	router.Route("/products", func(r chi.Router) {
		r.Get("/", HandlerProducts)
		r.Get("/products/{id}", HandlerProductsById)
		r.Get("/search/{priceGt}", HandlerSearchByPriceGt)
	})

	http.ListenAndServe(":8080", router)
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

func HandlerProducts(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(response).Encode(products); err != nil {
		http.Error(response, "Error products to json", http.StatusInternalServerError)
		return
	}
}

func HandlerProductsById(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		http.Error(response, "Invalid id", http.StatusBadRequest)
	}
	product, err := SearchById(id)

	if err != nil {
		http.Error(response, "Product not found", http.StatusNotFound)
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(response).Encode(product); err != nil {
		http.Error(response, "Error products to json", http.StatusInternalServerError)
		return
	}
}

func HandlerSearchByPriceGt(response http.ResponseWriter, request *http.Request) {
	priceGt, err := strconv.ParseFloat(chi.URLParam(request, "priceGt"), 2)
	if err != nil {
		http.Error(response, "Invalid price", http.StatusBadRequest)
	}
	productsUpThanPriceGt := SearchByPriceGt(priceGt)

	if err := json.NewEncoder(response).Encode(productsUpThanPriceGt); err != nil {
		http.Error(response, "Error products to json", http.StatusInternalServerError)
		return
	}

}

func SearchById(id int) (Product, error) {
	for _, product := range products {

		if product.Id == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Product not found")
}

func SearchByPriceGt(priceGt float64) []Product {

	var productUpThanPriceGt []Product

	for _, product := range products {
		if product.Price > priceGt {
			productUpThanPriceGt = append(productUpThanPriceGt, product)
		}
	}
	return productUpThanPriceGt

}
