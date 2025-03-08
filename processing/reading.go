package reading

import (
	"encoding/json"
	"net/http"
	erration "retsAPI/serv/error"

	"github.com/go-playground/validator/v10"
)

type BodyRequest struct {
	URL string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type BodyRespone struct {
	Alias string `json:"alias"`
	Error string `json:"error,omitempty"`
}

func ReadRequest() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var req BodyRequest
	
		err := json.NewDecoder(r.Body).Decode(&req)
		erration.LogError(err, "ERROR_REQUEST_DECODE")

		val := validator.New()

		err = val.Struct(req)
		erration.LogError(err, "ERROR_VALIDATION")
	}
}