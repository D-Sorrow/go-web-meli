package handlers

import (
	"encoding/json"
	"github.com/D-Sorrow/go-web-meli/clase-cinco-put-patch-delete/ejercicio-implement-put-patch-delete/internal/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: "error getting products",
		})
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(products); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: "error encoding products",
		})
		return
	}
}

func (product *Product) GetProductById(response http.ResponseWriter, request *http.Request) {

	productId, _ := strconv.Atoi(chi.URLParam(request, "id"))

	productById, err := product.productService.GetProductById(productId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(productById); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
	}

}
func (product *Product) GetProductByPriceGt(response http.ResponseWriter, request *http.Request) {

	productId, _ := strconv.ParseFloat(chi.URLParam(request, "priceGt"), 64)

	productByPriceGt, err := product.productService.GetProductByPriceGt(productId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(productByPriceGt); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
	}
}

func (product *Product) AddProduct(response http.ResponseWriter, request *http.Request) {

	var productAdd domain.Product

	if err := json.NewDecoder(request.Body).Decode(&productAdd); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
	}
	err := product.productService.AddProduct(productAdd)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
}

func (p *Product) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	var productUpdate domain.Product
	err := json.NewDecoder(request.Body).Decode(&productUpdate)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: "error decoder product",
		})
		return
	}
	err = p.productService.UpdateProduct(productUpdate)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
}

func (p *Product) PatchProduct(response http.ResponseWriter, request *http.Request) {
	idProduct, _ := strconv.Atoi(chi.URLParam(request, "id"))

	var attributes map[string]any

	if er := json.NewDecoder(request.Body).Decode(&attributes); er != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: "error decoder product patch",
		})
	}

	err := p.productService.PatchProduct(idProduct, attributes)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
		return
	}

	response.WriteHeader(http.StatusCreated)

}

func (p *Product) DeleteProduct(response http.ResponseWriter, request *http.Request) {
	idProduct, _ := strconv.Atoi(chi.URLParam(request, "id"))
	err := p.productService.DeleteProduct(idProduct)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(Message{
			Message: err.Error(),
		})
		return
	}
	response.WriteHeader(http.StatusCreated)
}
