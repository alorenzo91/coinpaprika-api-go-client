package coinpaprika

import (
	"encoding/json"
	"fmt"
)

type ChangeLogService service

// ChangeLog represents an ChangeLog.
type ChangeLog struct {
	CurrencyID *string `json:"currency_id"`
	OldID      *string `json:"old_id"`
	NewID      *string `json:"new_id"`
	ChangedAt  *string `json:"changed_at"`
}

// ChangeLogOptions specifies optional parameters for ChangeLog endpoints.
type ChangeLogOptions struct {
	Page int `url:"page,omitempty"`
}

// ChangeLog returns list of all ChangeLog listed on coinpaprika.
func (s *ChangeLogService) CoinsChangeLog(options *ChangeLogOptions) (changelog []*ChangeLog, err error) {
	url := fmt.Sprintf("%s/changelog/ids", baseURL)
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
