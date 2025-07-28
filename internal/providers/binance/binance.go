package binance

import (
	"encoding/json"

	"github.com/Dimashey/cro/internal/client/http"
)

type binance struct {
	client *http.Client
}

const binanceBaseURL = "https://api.binance.com/api/v3"

func New() *binance {
	client, _ := http.New(binanceBaseURL)

	return &binance{client}
}

func (b binance) Ticker() (BinanceTickerListResponse, error) {
	resp, err := b.client.Get("/ticker/24hr", nil)
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

func (b binance) PairTicker(pair string) (BinanceTickerPair, error) {
	resp, err := b.client.Get("/ticker/24hr", map[string]string{"symbol": pair})
	if err != nil {
		return BinanceTickerPair{}, err
	}

	defer resp.Body.Close()

	var ticker BinanceTickerPair
	err = json.NewDecoder(resp.Body).Decode(&ticker)
	if err != nil {
		return BinanceTickerPair{}, err
	}

	return ticker, nil
}

func (b binance) Candlesticks(options map[string]string) (BinanceCandlestickResponse, error) {
	resp, err := b.client.Get("/klines", options)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var candlesticks BinanceCandlestickResponse
	
	if err := json.NewDecoder(resp.Body).Decode(&candlesticks); err != nil {
		return nil, err
	}

	return candlesticks, nil
}
