package main

import (
	"errors"
	"net/http"

	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/domain/service"
	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/infrastucture/persistence/file"
	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/infrastucture/persistence/storage"
	"github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/transport/handlers"
	rout "github.com/D-Sorrow/go-web-meli/clase-seis-variables-entorno/ejercicios-variables-entorno/internal/transport/router"
	"github.com/go-chi/chi/v5"
)

func main() {

	storage := storage.NewStorageJSON("../../products.json")

	file, err := file.NewProductJSON(storage)

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
