package main

import (
	"errors"
	"github.com/D-Sorrow/go-web-meli/clase-cuatro-arquitectura-carpetas/ejercicio-uno-dominios/internal/domain/service"
	"github.com/D-Sorrow/go-web-meli/clase-cuatro-arquitectura-carpetas/ejercicio-uno-dominios/internal/infrastucture/persistence/file"
	"github.com/D-Sorrow/go-web-meli/clase-cuatro-arquitectura-carpetas/ejercicio-uno-dominios/internal/transport/handlers"
	rout "github.com/D-Sorrow/go-web-meli/clase-cuatro-arquitectura-carpetas/ejercicio-uno-dominios/internal/transport/router"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	file, err := file.NewProductJSON("products.json")

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

	http.ListenAndServe(":8083", routerPr)

}
