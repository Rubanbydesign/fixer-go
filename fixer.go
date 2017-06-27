// Package fixer is a client for the Fixer.io API
package fixer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
)

type Currency string

func (c Currency) String() string {
	return string(c)
}

type Rate struct {
	Currency Currency
	Rate     float64
}

type Rates []Rate

type apiResponse struct {
	Base  Currency `json:"base"`
	Date  string   `json:"date"`
	Rates map[Currency]float64
}

const (
	EUR Currency = "EUR"
	AUD Currency = "AUD"
	BGN Currency = "BGN"
	BRL Currency = "BRL"
	CAD Currency = "CAD"
	CHF Currency = "CHF"
	CNY Currency = "CNY"
	CZK Currency = "CZK"
	DKK Currency = "DKK"
	GBP Currency = "GBP"
	HKD Currency = "HKD"
	HRK Currency = "HRK"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	INR Currency = "INR"
	JPY Currency = "JPY"
	KRW Currency = "KRW"
	MXN Currency = "MXN"
	MYR Currency = "MYR"
	NOK Currency = "NOK"
	NZD Currency = "NZD"
	PHP Currency = "PHP"
	PLN Currency = "PLN"
	RON Currency = "RON"
	RUB Currency = "RUB"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	THB Currency = "THB"
	TRY Currency = "TRY"
	USD Currency = "USD"
	ZAR Currency = "ZAR"
)

func Convert(ctx context.Context, from Currency, to Currency, amt float64) (float64, error) {
	if from == to {
		return amt, nil
	}
	rates, err := Get(ctx, from)
	if err != nil {
		return 0, nil
	}
	for i := range rates {
		if rates[i].Currency != to {
			continue
		}
		return rates[i].Rate, nil
	}
	return 0, fmt.Errorf("Rates for %s not available", to)
}

func Get(ctx context.Context, base Currency) (Rates, error) {
	c := getClient(ctx)
	r, err := c.Get(fmt.Sprintf("https://api.fixer.io/latest?base=%s", base))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("Could not read response: %v", err)
		}
		return nil, fmt.Errorf("API error: %s", string(body))
	}
	var data apiResponse
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Could not decode response: %v", err)
	}
	i := 0
	rates := make(Rates, len(data.Rates))
	for key := range data.Rates {
		rates[i] = Rate{Currency: key, Rate: data.Rates[key]}
		i++
	}
	return rates, err
}

func Latest(ctx context.Context) (Rates, error) {
	return Get(ctx, EUR)
}
