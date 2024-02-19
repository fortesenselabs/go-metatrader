package models

type HistoricalDataResponse struct {
	// The symbol of the instrument
	Symbol string `json:"symbol"`
	// The time frame of the data
	TimeFrame string `json:"timeFrame"`
	// The action type of the data
	ActionType string `json:"actionType"`
	// The data
	Data []HistoricalData `json:"data"`
}

type HistoricalData struct {
	// The date of the data
	Date string `json:"date"`
	// The open price
	Open float64 `json:"open"`
	// The high price
	High float64 `json:"high"`
	// The low price
	Low float64 `json:"low"`
	// The close price
	Close float64 `json:"close"`
	// The volume
	Volume float64 `json:"volume"`
}

// "datetime",
// "open",
// "high",
// "low",
// "close",
// "tick_volume",
// "real_volume",
// "spread",
