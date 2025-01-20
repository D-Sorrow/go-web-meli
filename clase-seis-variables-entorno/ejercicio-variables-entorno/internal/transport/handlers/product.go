package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/domain"
	res "github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/plataform/web/response"
	"github.com/go-chi/chi/v5"
)

type Product struct {
	productService domain.ProductService
}

type Message struct {
	Message string `json:"message"`
}

func Router(hd *Product) *chi.Mux {

	rt := chi.NewRouter()
	rt.Get("/", hd.GetProducts)
	rt.Get("/{id}", hd.GetProductById)
	rt.Get("/search/{priceGt}", hd.GetProductByPriceGt)
	rt.Post("/", hd.AddProduct)
	rt.Put("/", hd.UpdateProduct)
	rt.Patch("/{id}", hd.PatchProduct)
	rt.Delete("/{id}", hd.DeleteProduct)
	return rt
}

func NewHandlerProduct(productService domain.ProductService) *Product {
	return &Product{productService: productService}
}

func (product *Product) GetProducts(response http.ResponseWriter, request *http.Request) {
	products, err := product.productService.GetProducts()
	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "Error get products to server")
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(products); err != nil {
		res.SetError(response, http.StatusInternalServerError, "Error json encoder")
		return
	}
}

func (product *Product) GetProductById(response http.ResponseWriter, request *http.Request) {

	productId, _ := strconv.Atoi(chi.URLParam(request, "id"))

	productById, err := product.productService.GetProductById(productId)

	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "Error get product by id")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(productById); err != nil {
		res.SetError(response, http.StatusInternalServerError, "Error json encoder")
		return
	}

}
func (product *Product) GetProductByPriceGt(response http.ResponseWriter, request *http.Request) {
	productId, _ := strconv.ParseFloat(chi.URLParam(request, "priceGt"), 64)

	productByPriceGt, err := product.productService.GetProductByPriceGt(productId)

	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "Error get product by price")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(productByPriceGt); err != nil {
		res.SetError(response, http.StatusInternalServerError, "Error json encoder")
		return
	}
}

func (product *Product) AddProduct(response http.ResponseWriter, request *http.Request) {

	token := request.Header.Get("TOKEN_API")
	if token != os.Getenv("TOKEN_API") || token == "" {
		res.SetError(response, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var productAdd domain.Product

	if err := json.NewDecoder(request.Body).Decode(&productAdd); err != nil {
		res.SetError(response, http.StatusInternalServerError, "error json decoder")
		return
	}
	err := product.productService.AddProduct(productAdd)
	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "error add product")
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
}

func (p *Product) UpdateProduct(response http.ResponseWriter, request *http.Request) {

	token := request.Header.Get("TOKEN_API")
	if token != os.Getenv("TOKEN_API") || token == "" {
		res.SetError(response, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var productUpdate domain.Product
	err := json.NewDecoder(request.Body).Decode(&productUpdate)
	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "error decode json")
		return
	}
	err = p.productService.UpdateProduct(productUpdate)
	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "error update product")
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
}

func (p *Product) PatchProduct(response http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("TOKEN_API")
	if token != os.Getenv("TOKEN_API") || token == "" {
		res.SetError(response, http.StatusUnauthorized, "Unauthorized")
		return
	}
	idProduct, _ := strconv.Atoi(chi.URLParam(request, "id"))

	var attributes map[string]any

	if er := json.NewDecoder(request.Body).Decode(&attributes); er != nil {
		res.SetError(response, http.StatusInternalServerError, "error decode json")
		return
	}

	err := p.productService.PatchProduct(idProduct, attributes)

	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "error update product")
		return
	}

	response.WriteHeader(http.StatusNoContent)

}

func (p *Product) DeleteProduct(response http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("TOKEN_API")
	if token != os.Getenv("TOKEN_API") || token == "" {
		res.SetError(response, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idProduct, _ := strconv.Atoi(chi.URLParam(request, "id"))
	err := p.productService.DeleteProduct(idProduct)
	if err != nil {
		res.SetError(response, http.StatusInternalServerError, "error delete product")
		return
	}
	response.WriteHeader(http.StatusOK)
}
