package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// ContractsService is used for getting contract information.
// API Documentation: https://api.coinpaprika.com/#tag/Contracts
type ContractsService service

type ContractsOptions struct {
}

type ContractsResponse struct {
	Address string `json:"address"`
	ID      string `json:"id"`
	Type    string `json:"type"`
}

// List retrieves a list of contract addresses from the Coinpaprika API.
// It constructs the URL for the request, sends a GET request to the API,
// and unmarshals the response body into a slice of contract address strings.
//
// Returns a slice of contract address strings and an error, if any.
func (t *ContractsService) List() (contracts []*string, err error) {
	url := baseURL + "/contracts"
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

// GetContractaddressessForPlatform retrieves the contract addresses for a given platform.
// It takes a platformID as a parameter and returns a slice of ContractsResponse and an error.
//
// Parameters:
//   - platformID: The ID of the platform for which to retrieve contract addresses.
//
// Returns:
//   - contracts: A slice of ContractsResponse containing the contract addresses.
//   - err: An error if the request fails or the response cannot be unmarshaled.
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

// GetTickerByContractAddress retrieves the ticker information for a given contract address on a specified platform.
//
// Parameters:
//   - platformID: The ID of the platform (e.g., ethereum).
//   - contractAddress: The contract address of the token.
//
// Returns:
//   - ticker: A pointer to the Ticker struct containing the ticker information.
//   - err: An error if the request fails or the response cannot be unmarshaled.
func (t *ContractsService) GetTickerByContractAddress(platformID, contractAddress string) (ticker *Ticker, err error) {
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
