package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Fprintf(writer, "pong")
	default:
		fmt.Errorf("Method not supported")
	}
}
