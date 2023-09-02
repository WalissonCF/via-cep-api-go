package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

const viaCEPURL = "https://viacep.com.br/ws/"

func main() {
	cep := "06341420"

	dadosCEP, err := ConsultaCEP(cep)
	if err != nil {
		fmt.Println("Erro ao consultar CEP:", err)
		return
	}

	fmt.Println("Dados do CEP:", cep)
	fmt.Println(dadosCEP)
}

func ConsultaCEP(cep string) (ViaCEPResponse, error) {
	url := viaCEPURL + cep + "/json/"

	var viaCEPResponse ViaCEPResponse

	response, err := http.Get(url)
	if err != nil {
		return viaCEPResponse, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return viaCEPResponse, err
	}

	err = json.Unmarshal(body, &viaCEPResponse)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return viaCEPResponse, err
	}

	return viaCEPResponse, nil
}
