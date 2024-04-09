package models

type Order struct {
	Id         string  `json:"id"`
	Magic      int     `json:"magic"`
	Symbol     string  `json:"symbol"`
	Type       string  `json:"type"`
	TimeSetup  int     `json:"time_setup"`
	Open       float64 `json:"open"`
	Stoploss   float64 `json:"stoploss"`
	Takeprofit float64 `json:"takeprofit"`
	Volume     float64 `json:"volume"`
}

type Orders []Order
