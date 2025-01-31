package handlers_test

import (
	"bytes"
	"errors"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain/service"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProduct_GetProducts(t *testing.T) {
	mockService := &service.ProductMock{
		GetProductsFunc: func() ([]domain.Product, error) {
			return []domain.Product{
				{
					Id:           1,
					Name:         "Canned, Rings",
					Quantity:     345,
					Code_value:   "SNSNQ23",
					Is_published: true,
					Expiration:   "09/08/2021",
					Price:        352.79,
				},
				{
					Id:           2,
					Name:         "Margarinrre",
					Quantity:     439,
					Code_value:   "8D",
					Is_published: true,
					Expiration:   "15/12/2029",
					Price:        71.42,
				},
			}, nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Get("/products", handler.GetProducts)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
	expectedBody := `[{"id":1,"name":"Canned, Rings","quantity":345,"code_value":"SNSNQ23","is_published":true,"expiration":"09/08/2021","price":352.79},{"id":2,"name":"Margarinrre","quantity":439,"code_value":"8D","is_published":true,"expiration":"15/12/2029","price":71.42}]`

	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, expectedBody, rec.Body.String())
	require.Equal(t, expectedHeader, rec.Header())

}

func TestProduct_GetProductById(t *testing.T) {
	mockService := &service.ProductMock{
		GetProductByIDFunc: func(id int) (domain.Product, error) {
			return domain.Product{
				Id:           1,
				Name:         "Canned, Rings",
				Quantity:     345,
				Code_value:   "SNSNQ23",
				Is_published: true,
				Expiration:   "09/08/2021",
				Price:        352.79,
			}, nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()

	router.Get("/{id}", handler.GetProductById)

	req := httptest.NewRequest(http.MethodGet, "/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	expectedBody := `{"id":1,"name":"Canned, Rings","quantity":345,"code_value":"SNSNQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`

	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, expectedBody, rec.Body.String())
	require.Equal(t, expectedHeader, rec.Header())

}

func TestProduct_CreateProduct(t *testing.T) {
	mockService := &service.ProductMock{
		AddProductFunc: func(product domain.Product) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Post("/products", handler.AddProduct)

	body := []byte(`{"name":"Canned, Rings","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "123456")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	require.Equal(t, http.StatusCreated, rec.Code)
	require.Equal(t, expectedHeader, rec.Header())
}

func TestProduct_DeleteProduct(t *testing.T) {
	mockService := &service.ProductMock{
		DeleteProductFunc: func(id int) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()
	router.Delete("/products/{id}", handler.DeleteProduct)

	req := httptest.NewRequest(http.MethodDelete, "/products/2", nil)
	req.Header.Add("TOKEN_API", "123456")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNoContent, rec.Code)
}

func TestProduct_GetProductById_ErrId(t *testing.T) {

	mockService := &service.ProductMock{
		GetProductByIDFunc: func(id int) (domain.Product, error) {
			return domain.Product{}, errors.New("Error id is not valid")
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()
	router.Get("/products/{id}", handler.GetProductById)

	req := httptest.NewRequest(http.MethodGet, "/products/1A", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestProduct_UpdatePutProduct_ErrId(t *testing.T) {
	mockService := &service.ProductMock{
		UpdateProductFun: func(product domain.Product) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Put("/product", handler.UpdateProduct)

	body := []byte(`{"id": 1A,"name":"CannedRings","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPut, "/product", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "123456")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestProduct_UpdatePatchProduct_ErrId(t *testing.T) {
	mockService := &service.ProductMock{
		PatchProductFunc: func(id int, att map[string]any) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Patch("/product/{id}", handler.PatchProduct)

	body := []byte(`"name":"Canned","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPatch, "/product/A1", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "123456")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestProduct_DeleteProduct_ErrId(t *testing.T) {
	mockService := &service.ProductMock{
		DeleteProductFunc: func(id int) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()
	router.Delete("/products/{id}", handler.DeleteProduct)

	req := httptest.NewRequest(http.MethodDelete, "/products/2A", nil)
	req.Header.Add("TOKEN_API", "123456")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestProduct_ProductGetById_NotFound(t *testing.T) {
	mockService := &service.ProductMock{
		GetProductByIDFunc: func(id int) (domain.Product, error) {
			return domain.Product{}, errors.New("Product not found")
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()
	router.Get("/products/{id}", handler.GetProductById)

	req := httptest.NewRequest(http.MethodGet, "/products/200", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestProduct_ProductPatch_NotFound(t *testing.T) {
	mockService := &service.ProductMock{
		PatchProductFunc: func(id int, att map[string]any) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Patch("/product/{id}", handler.PatchProduct)

	body := []byte(`"name":"Canned","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPatch, "/product/199", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "123456")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestProduct_DeleteProduct_NotFound(t *testing.T) {
	mockService := &service.ProductMock{
		DeleteProductFunc: func(id int) error {
			return errors.New("Product not found")
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()
	router.Delete("/products/{id}", handler.DeleteProduct)

	req := httptest.NewRequest(http.MethodDelete, "/products/2000", nil)
	req.Header.Add("TOKEN_API", "123456")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestProduct_CreateProduct_Unauthorized(t *testing.T) {
	mockService := &service.ProductMock{
		AddProductFunc: func(product domain.Product) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Post("/products", handler.AddProduct)

	body := []byte(`{"name":"Canned, Rings","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "12342356")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
func TestProduct_UpdatePutProduct_Unauthorized(t *testing.T) {
	mockService := &service.ProductMock{
		UpdateProductFun: func(product domain.Product) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Put("/product", handler.UpdateProduct)

	body := []byte(`{"id": 1A,"name":"CannedRings","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPut, "/product", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "1232323456")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
func TestProduct_UpdatePatchProduct_Unauthorized(t *testing.T) {
	mockService := &service.ProductMock{
		PatchProductFunc: func(id int, att map[string]any) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)

	router := chi.NewRouter()
	router.Patch("/product/{id}", handler.PatchProduct)

	body := []byte(`"name":"Canned","quantity":345,"code_value":"SQ23","is_published":true,"expiration":"09/08/2021","price":352.79}`)

	req := httptest.NewRequest(http.MethodPatch, "/product/A1", bytes.NewBuffer(body))
	req.Header.Add("TOKEN_API", "123434456")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestProduct_DeleteProduct_Unauthorized(t *testing.T) {
	mockService := &service.ProductMock{
		DeleteProductFunc: func(id int) error {
			return nil
		},
	}

	handler := handlers.NewHandlerProduct(mockService)
	router := chi.NewRouter()
	router.Delete("/products/{id}", handler.DeleteProduct)

	req := httptest.NewRequest(http.MethodDelete, "/products/2A", nil)
	req.Header.Add("TOKEN_API", "12342356")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
