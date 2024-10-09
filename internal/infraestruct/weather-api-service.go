package infraestruct

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

type WeatherReturn struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherApiService struct {
	apiKey string
}

func NewWeatherServiceWeatherApi(apiKey string) *WeatherApiService {
	return &WeatherApiService{
		apiKey: apiKey,
	}
}

func (w *WeatherApiService) DiscoverWeather(location string) (float64, error) {
	encodedLocation := url.QueryEscape(location)
	response, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + w.apiKey + "&q=" + encodedLocation)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("error status code %d from Wheater API", response.StatusCode)
		return 0.00, errors.New("error occurred while processing your request of Wheater API response")
	}

	if err != nil {
		return 0.00, err
	}

	body, err := io.ReadAll(response.Body)

	var WeatherReturn *WeatherReturn

	errJson := json.Unmarshal(body, &WeatherReturn)
	if errJson != nil {
		return 0.00, errJson
	}

	return WeatherReturn.Current.TempC, nil
}
