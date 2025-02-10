package coinpaprika

import (
	"encoding/json"
)

// SearchService is used for search requests
// API Documentation: https://api.coinpaprika.com/#tag/Search
type SearchService service

// SearchResult represents a result of search API endpoint.
type SearchResult struct {
	Currencies []*Coin     `json:"currencies"`
	ICOS       []*ICO      `json:"icos"`
	Exchanges  []*Exchange `json:"exchanges"`
	People     []*Person   `json:"people"`
	Tags       []*Tag      `json:"tags"`
}

// SearchOptions specifies optional parameters for search endpoint.
type SearchOptions struct {
	Query string `url:"q"`

	// Comma separated categories to include in search results.
	// Available options: currencies|exchanges|icos|people|tags.
	// Eg. "currencies,exchanges"
	Categories string `url:"c,omitempty"`

	// The number of results per category.
	Limit int `url:"limit,omitempty"`

	Modifier string `url:"modifier,omitempty"`
}

// Search returns a list of currencies, exchanges, icos, people and tags for given query.
func (s *SearchService) Search(options *SearchOptions) (searchResult *SearchResult, err error) {
	url := baseURL + "/search"
	url, err = constructURL(url, options)

	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &searchResult)

	return searchResult, err
}
