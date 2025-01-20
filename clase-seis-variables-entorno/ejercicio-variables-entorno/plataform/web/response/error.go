package response

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string
	code    int
}

func SetError(response http.ResponseWriter, statusCode int, message string) {

	err := errorResponse{}

	if statusCode > 299 && statusCode < 600 {
		err.Message = message
		err.code = statusCode
	}
	body, errorJson := json.Marshal(err)

	if errorJson != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.WriteHeader(err.code)
	response.Header().Set("Content-type", "application/json")
	response.Write(body)

}
