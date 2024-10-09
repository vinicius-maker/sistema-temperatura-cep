package usecase

import (
	"errors"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/entity"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/service"
)

var ErrCepCodeNotFound = errors.New("can not find zipcode")

type DiscoverWeatherByLocationDTO struct {
	CepCode string `json:"zip_code"`
}

type DiscoverWeatherByLocationOutputDTO struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type DiscoverWeatherByLocation struct {
	viaCepService  service.ViaCepService
	weatherService service.WeatherService
}

func NewDiscoverWeatherByLocation(viaCepService service.ViaCepService, weatherService service.WeatherService) *DiscoverWeatherByLocation {
	return &DiscoverWeatherByLocation{
		viaCepService:  viaCepService,
		weatherService: weatherService,
	}
}

func (d *DiscoverWeatherByLocation) Execute(inputDTO DiscoverWeatherByLocationDTO) (DiscoverWeatherByLocationOutputDTO, error) {
	outputDTO := DiscoverWeatherByLocationOutputDTO{}

	cepCode, err := entity.NewCepCode(inputDTO.CepCode)

	if err != nil {
		return outputDTO, err
	}

	location, err := d.viaCepService.SearchCep(cepCode)
	if err != nil {
		return outputDTO, err
	}

	if location == "" {
		return outputDTO, ErrCepCodeNotFound
	}

	weather, err := d.weatherService.DiscoverWeather(location)
	if err != nil {
		return outputDTO, err
	}

	converter := entity.WeatherConverter{
		Celsius: weather,
	}

	outputDTO.TempC = weather
	outputDTO.TempF = converter.ToFahrenheit()
	outputDTO.TempK = converter.ToKelvin()

	return outputDTO, err
}
