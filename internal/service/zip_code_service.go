package service

import (
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/entity"
)

type ViaCepService interface {
	SearchCep(zipCode *entity.CepCode) (string, error)
}
