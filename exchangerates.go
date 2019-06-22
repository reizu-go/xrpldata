package xrplda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// GetExchangeRatesOptions contains optional parameters for the GetExchangeRates
// method.GetExchangeRatesOptions
type GetExchangeRatesOptions struct {
	Date   time.Time
	Strict bool
}

type exchangeRatesResponse struct {
	Rate string
}

// GetExchangeRates retrieves an exchange rate for a given currency pair at a
// specific time.
// https://developers.ripple.com/data-api.html#get-exchange-rates
func (c *Client) GetExchangeRates(base, counter string, opts *GetExchangeRatesOptions) (string, *Response, error) {
	path := fmt.Sprintf("/exchange_rates/%s/%s", base, counter)

	if opts != nil {
		v := url.Values{}

		if !opts.Date.IsZero() {
			v.Set("date", FormatTime(opts.Date))
		}

		if opts.Strict {
			v.Set("strict", "true")
		}

		path += "?" + v.Encode()
	}

	res, err := c.Do(http.MethodGet, path)

	data := &exchangeRatesResponse{}
	err = json.Unmarshal(res.Body, data)
	if err != nil {
		return "", res, err
	}

	return data.Rate, res, nil
}