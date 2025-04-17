package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://viacep.com.br/ws/%s/json"
)

var (
	ErrCepNotFound    = errors.New("not_found_error")
	errResponseDecode = errors.New("response_decode_error")
)

type APIResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Client struct {
	cli *http.Client
}

func New(cli http.Client) *Client {
	return &Client{
		cli: &cli,
	}
}

func (c *Client) FetchLocation(cep string) (response APIResponse, err error) {
	resp, err := c.cli.Get(fmt.Sprintf(baseURL, cep))
	if err != nil {
		return response, errors.New("error making viacep request")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, errResponseDecode
	}
	if resp.StatusCode == http.StatusNotFound {
		return response, ErrCepNotFound
	}
	return response, nil
}
