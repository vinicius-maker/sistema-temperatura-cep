package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/controller"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/infraestruct"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("not found environment variable API_KEY")
	}

	viaCepService := infraestruct.NewServiceViaCep()
	weatherService := infraestruct.NewWeatherServiceWeatherApi(apiKey)

	app := usecase.NewDiscoverWeatherByLocation(viaCepService, weatherService)

	weatherController := controller.NewWeatherController(app)

	http.HandleFunc("/discover-temperature", weatherController.Handle)

	http.ListenAndServe(":8080", nil)
}
