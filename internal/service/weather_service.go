package service

type WeatherService interface {
	DiscoverWeather(location string) (float64, error)
}
