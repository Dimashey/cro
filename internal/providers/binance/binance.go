package binance

import "encoding/json"
import "github.com/Dimashey/cro/internal/client/http"

type binance struct{
	client *http.Client
}

const binanceBaseURL = "https://api.binance.com/api/v3"

func New() *binance {
	client, _ := http.New(binanceBaseURL)

	return &binance{client}
}

func (b binance) Ticker() (BinanceTickerListResponse, error) {
	resp, err := b.client.Get("/ticker/24hr")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var ticker BinanceTickerListResponse
	err = json.NewDecoder(resp.Body).Decode(&ticker)

	if err != nil {
		return nil, err
	}

	return ticker, nil
}
