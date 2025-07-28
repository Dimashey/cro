package binance

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type BinanceCandlestick struct {
	OpenTime   int
	OpenPrice  float64
	HighPrice  float64
	LowPrice   float64
	ClosePrice float64
	CloseTime  int
}

func getFloatFromData(data any) (float64, error) {
	value, ok := data.(string)

	if !ok {
		return 0, fmt.Errorf("invalid value")
	}

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value")
	}

	return v, nil
}

func getIntFromData(data any) (int, error) {
	value, ok := data.(float64)

	if !ok {
		return 0, fmt.Errorf("invalid value")
	}

	return int(value), nil
}

func (bc *BinanceCandlestick) UnmarshalJSON(data []byte) error {
	var tuple []interface{}

	if err := json.Unmarshal(data, &tuple); err != nil {
		return err
	}

	openTime, err := getIntFromData(tuple[0])
	if err != nil {
		fmt.Println("T")
		return err
	}

	openPrice, err := getFloatFromData(tuple[1])
	if err != nil {
		return err
	}

	highPrice, err := getFloatFromData(tuple[2])
	if err != nil {
		return err
	}

	lowPrice, err := getFloatFromData(tuple[3])
	if err != nil {
		return err
	}

	closePrice, err := getFloatFromData(tuple[4])
	if err != nil {
		return err
	}

	closeTime, err := getIntFromData(tuple[6])
	if err != nil {
		return err
	}

	bc.OpenTime = openTime
	bc.OpenPrice = openPrice
	bc.HighPrice = highPrice
	bc.LowPrice = lowPrice
	bc.ClosePrice = closePrice
	bc.CloseTime = closeTime

	period := (bc.CloseTime - bc.OpenTime) / 3600

	fmt.Println(period)

	return nil
}
