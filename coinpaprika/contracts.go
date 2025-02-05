package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// ContractsService is used for getting contract information.
type ContractsService service

type ContractsOptions struct {
}

type ContractsResponse struct {
	Address string `json:"address"`
	ID      string `json:"id"`
	Type    string `json:"type"`
}

func (t *ContractsService) List() (contracts []*string, err error) {
	url := fmt.Sprintf("%s/contracts", baseURL)
	url, err = constructURL(url, nil)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(t.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &contracts)
	return contracts, err

}

func (t *ContractsService) GetContractaddressessForPlatform(platformID string) (contracts []*ContractsResponse, err error) {
	url := fmt.Sprintf("%s/contracts/%s", baseURL, platformID)
	url, err = constructURL(url, nil)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(t.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &contracts)
	return contracts, err
}

func (t *ContractsService) GetTickerByContractAddress(platformID string, contractAddress string) (ticker *Ticker, err error) {
	url := fmt.Sprintf("%s/contracts/%s/%s", baseURL, platformID, contractAddress)
	url, err = constructURL(url, nil)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(t.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ticker)
	return ticker, err
}
