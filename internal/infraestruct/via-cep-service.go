package infraestruct

import (
	"encoding/json"
	"errors"
	"github.com/vinicius-maker/sistema-temperatura-cep/internal/entity"
	"io"
	"log"
	"net/http"
)

type ViaCepStruct struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

var ErrRequestCepCode = errors.New("error occurred while processing your request of zipcode")

type ServiceViaCep struct {
}

func NewServiceViaCep() *ServiceViaCep {
	return &ServiceViaCep{}
}

func (c ServiceViaCep) SearchCep(cepCode *entity.CepCode) (string, error) {
	response, err := http.Get("http://viacep.com.br/ws/" + cepCode.CepCode + "/json/")
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("error status code %d from ViaCEP", response.StatusCode)
		return "", ErrRequestCepCode
	}

	if err != nil {
		log.Printf("Error: %v", err)
		return "", ErrRequestCepCode
	}

	body, err := io.ReadAll(response.Body)

	var viaCepStruct ViaCepStruct
	err = json.Unmarshal(body, &viaCepStruct)

	return viaCepStruct.Localidade, err
}
