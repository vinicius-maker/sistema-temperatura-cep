package controller

import (
	"encoding/json"
	"errors"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/entity"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/usecase"
	"net/http"
)

type WeatherController struct {
	usecase *usecase.DiscoverWeatherByLocation
}

func NewWeatherController(usecase *usecase.DiscoverWeatherByLocation) WeatherController {
	return WeatherController{
		usecase: usecase,
	}
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (wc WeatherController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cepCode := r.URL.Query().Get("cep")

	var dto usecase.DiscoverWeatherByLocationDTO
	dto.CepCode = cepCode

	output, err := wc.usecase.Execute(dto)

	if err != nil {
		switch {
		case errors.Is(err, entity.ErrInvalidCepCode):
			writeErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		case errors.Is(err, usecase.ErrCepCodeNotFound):
			writeErrorResponse(w, http.StatusNotFound, err.Error())
			return
		default:
			writeErrorResponse(w, http.StatusInternalServerError, "error processing the request")
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(output)
}
