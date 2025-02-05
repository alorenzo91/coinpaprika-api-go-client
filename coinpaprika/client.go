package coinpaprika

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	userAgent   = "Coinpaprika API Client - Go"
	baseFreeURL = "https://api.coinpaprika.com/v1"
	baseProURL  = "https://api-pro.coinpaprika.com/v1"
)

var baseURL = baseFreeURL

// Client: Is a struct that contains a single instance of a http client and all the
// services that can be used to interact with the Coin Paprika API(https://api.coinpaprika.com/).
type (
	Client struct {
		httpClient     *http.Client
		Tickers        TickersService        // TickersService is used for getting ticker information.
		Search         SearchService         // SearchService is used for search requests.
		PriceConverter PriceConverterService // PriceConverterService is used for price converter requests.
		Coins          CoinsService          // CoinsService is used for getting coin information.
		Global         GlobalService         // GlobalService is used for getting global information.
		Tags           TagsService           // TagsService is used for getting tags information.
		People         PeopleService         // PeopleService is used for getting people information.
		Exchanges      ExchangesService      // ExchangesService is used for getting exchange information.
		Contracts      ContractsService      // ContractsService is used for getting contract information.
	}
)

type (
	service struct {
		httpClient *http.Client
	}
)

type (
	authTransport struct {
		baseTransport http.RoundTripper
		apiKey        string
	}
)

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", t.apiKey)
	return t.baseTransport.RoundTrip(req)
}

type ClientOptions func(a *Client)

// NewClient creates a new client to work with coinpaprika API.
// and injects the client into each service. This allows the client to be used with a custom http client.
func NewClient(httpClient *http.Client, opts ...ClientOptions) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Tickers.httpClient = c.httpClient
	c.Search.httpClient = c.httpClient
	c.PriceConverter.httpClient = c.httpClient
	c.Coins.httpClient = c.httpClient
	c.Global.httpClient = c.httpClient
	c.Tags.httpClient = c.httpClient
	c.People.httpClient = c.httpClient
	c.Exchanges.httpClient = c.httpClient
	c.Contracts.httpClient = c.httpClient

	return c
}

// WithAPIKey sets API key enabling access to Coinpaprika Pro API.
// https://api-pro.coinpaprika.com is used.
func WithAPIKey(apiKey string) ClientOptions {
	return func(a *Client) {
		baseURL = baseProURL

		baseTransport := http.DefaultTransport
		if a.httpClient.Transport != nil {
			baseTransport = a.httpClient.Transport
		}

		a.httpClient.Transport = &authTransport{
			apiKey:        apiKey,
			baseTransport: baseTransport,
		}
	}
}

// constructURL constructs a URL with query parameters from the given options.
// If the options parameter is a nil pointer, it returns the rawURL as is.
// It parses the rawURL and encodes the options as query parameters, appending them to the URL.
//
// Parameters:
//   - rawURL: The base URL to which query parameters will be added.
//   - options: An interface containing the query parameters to be added to the URL.
//
// Returns:
//   - A string representing the constructed URL with query parameters.
//   - An error if there is an issue parsing the URL or encoding the query parameters.
func constructURL(rawURL string, options interface{}) (string, error) {
	if v := reflect.ValueOf(options); v.Kind() == reflect.Ptr && v.IsNil() {
		return rawURL, nil
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return rawURL, err
	}

	values, err := query.Values(options)
	if err != nil {
		return rawURL, err
	}

	parsedURL.RawQuery = values.Encode()

	return parsedURL.String(), nil
}

// sendGET sends an HTTP GET request to the specified URL using the provided HTTP client.
// It returns the response body as a byte slice if the request is successful, or an error otherwise.
//
// Parameters:
//   - client: The HTTP client to use for sending the request.
//   - url: The URL to send the GET request to.
//
// Returns:
//   - []byte: The response body as a byte slice.
//   - error: An error if the request fails or the response status code is not 200 OK.
func sendGET(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil) //nolint:gocritic
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v, body: %s", response.StatusCode, string(body))
	}

	return body, nil
}
