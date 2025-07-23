package binance

type BinanceTickerPair struct {
	Symbol    string `json:"symbol"`
	Volume    string `json:"volume"`
	LastPrice string `json:"lastPrice"`
}

type BinanceTickerListResponse []BinanceTickerPair
