package coinpaprika

import (
	"encoding/json"
)

// ChangeLogService provides methods for interacting with change log-related endpoints.
type ChangeLogService service

// ChangeLog represents a change log entry for a currency.
// It contains details about the currency ID change, including the old and new IDs, and the timestamp of the change.
type ChangeLog struct {
	CurrencyID *string `json:"currency_id"` // CurrencyID: The ID of the currency that was changed.
	OldID      *string `json:"old_id"`      // OldID: The previous ID of the currency.
	NewID      *string `json:"new_id"`      // NewID: The new ID of the currency after the change.
	ChangedAt  *string `json:"changed_at"`  // ChangedAt: The timestamp when the change occurred.
}

// ChangeLogOptions specifies optional parameters for querying change log endpoints.
// It allows pagination and other filtering options.
//
// Fields:
//   - Page: The page number to retrieve for paginated results (optional).
type ChangeLogOptions struct {
	Page int `url:"page,omitempty"`
}

// CoinsChangeLog retrieves the change log for coins based on the provided options.
// It takes a ChangeLogOptions parameter and returns a slice of ChangeLog pointers and an error.
//
// Parameters:
//   - options: The options to filter or customize the change log retrieval.
//
// Returns:
//   - changelog: A slice of ChangeLog pointers containing the change log data.
//   - err: An error if the request fails, the response cannot be unmarshaled, or the URL construction fails.
func (s *ChangeLogService) CoinsChangeLog(options *ChangeLogOptions) (changelog []*ChangeLog, err error) {
	url := baseURL + "/changelog/ids"
	url, err = constructURL(url, options)

	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &changelog)

	return changelog, err
}
