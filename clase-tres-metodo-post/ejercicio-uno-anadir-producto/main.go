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
	"time"
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
type Message struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
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
		r.Get("/{id}", HandlerProductsById)
		r.Get("/search/{priceGt}", HandlerSearchByPriceGt)
		r.Post("/", AddProduct)
	})

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		return
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

func HandlerProducts(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(products); err != nil {
		http.Error(response, "Error products to json", http.StatusInternalServerError)
		return
	}
}

func HandlerProductsById(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	response.Header().Set("Content-Type", "application/json")

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Message{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	product, err := SearchById(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Message{
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(product); err != nil {
		json.NewEncoder(response).Encode(product)
		return
	}
}

func HandlerSearchByPriceGt(response http.ResponseWriter, request *http.Request) {
	priceGt, err := strconv.ParseFloat(chi.URLParam(request, "priceGt"), 2)
	if err != nil {
		http.Error(response, "Invalid price", http.StatusBadRequest)
	}

	productsUpThanPriceGt := SearchByPriceGt(priceGt)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(productsUpThanPriceGt); err != nil {
		http.Error(response, "Error products to json", http.StatusInternalServerError)
		return
	}

}

func AddProduct(response http.ResponseWriter, request *http.Request) {
	var newProduct Product
	response.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(request.Body).Decode(&newProduct); err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(products) > 0 {
		newProduct.Id = (products)[len(products)-1].Id + 1
	} else {
		newProduct.Id = 1
	}

	if ValidateCodeValue(newProduct.Code_value) {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Message{
			Error:   true,
			Message: "Error code already exits",
		})
		return
	}

	if ValidateDateExpiration(newProduct.Expiration) {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Message{
			Error:   true,
			Message: "Error expiration",
		})
		return
	}

	products = append(products, newProduct)

	response.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(response).Encode(newProduct); err != nil {
		http.Error(response, "Error products to json", http.StatusInternalServerError)
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

func ValidateCodeValue(code string) bool {
	for _, product := range products {
		if product.Code_value == code {
			return true
		}
	}
	return false
}

func ValidateDateExpiration(date string) bool {
	dateExp, err := time.Parse("2006/01/02", date)
	if err != nil {
		return true
	}
	if dateExp.Year() < time.Now().Year() {
		return true
	}
	return false
}
