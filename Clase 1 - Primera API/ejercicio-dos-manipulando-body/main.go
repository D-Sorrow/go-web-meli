package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	http.HandleFunc("/greetings", greetingHandler)
	http.ListenAndServe(":8080", nil)
}

func greetingHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "No se pudo leer el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		var person Person
		err = json.Unmarshal(body, &person)
		if err != nil {
			http.Error(writer, "JSON inválido", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(writer, "Hello %s %s", person.FirstName, person.LastName)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(writer, "Method not supported")
	}
}
