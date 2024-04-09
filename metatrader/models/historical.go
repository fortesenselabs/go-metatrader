package models

// type HistoricalDataResponse struct {
// 	// The symbol of the instrument
// 	Symbol string `json:"symbol"`
// 	// The time frame of the data
// 	TimeFrame string `json:"timeFrame"`
// 	// The action type of the data
// 	ActionType string `json:"actionType"`
// 	// The data
// 	Rates []Rate `json:"rates"`
// }

type Rate struct {
	// Date       time.Time `json:"datetime"`
	Date       string  `json:"datetime"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	TickVolume float64 `json:"tick_volume"`
	RealVolume float64 `json:"real_volume"`
	Spread     int64   `json:"spread"`
}

// "datetime",
// "open",
// "high",
// "low",
// "close",
// "tick_volume",
// "real_volume",
// "spread",

type Tick struct {
	Timestamp string  `json:"timestamp"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
}

type Trade struct {
	Ticket    string `json:"ticket"`
	Timestamp string `json:"timestamp"`
	Price     string `json:"high"`
	Volume    string `json:"low"`
	Symbol    string `json:"close"`
	Type      string `json:"tick_volume"`
	Entry     string `json:"real_volume"`
	Profit    string `json:"spread"`
}
