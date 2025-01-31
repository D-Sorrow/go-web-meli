package main

import (
	"errors"
	"net/http"

	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/domain/service"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/infrastucture/persistence/file"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/infrastucture/persistence/storage"
	"github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/transport/handlers"
	rout "github.com/D-Sorrow/go-web-meli/clase-siete-test/ejercicio-test/internal/transport/router"
	"github.com/go-chi/chi/v5"
)

func main() {

	storage := storage.NewStorageJSON("products.json")

	db, _ := storage.ReadAll()

	file, err := file.NewProductJSON(db)

	if err != nil {
		errors.New(err.Error())
	}

	productService := service.NewProductService(file)

	handler := handlers.NewHandlerProduct(productService)

	router := handlers.Router(handler)

	mapRut := map[string]*chi.Mux{
		"/product": router,
	}

	routerPr := rout.NewRouter(mapRut)

	http.ListenAndServe(":8081", routerPr)

}
