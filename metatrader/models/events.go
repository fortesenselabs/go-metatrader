package models

type TickEvent struct {
	Event string        `json:"event"`
	Data  TickEventData `json:"data"`
}

type TickEventData struct {
	Symbol    string        `json:"symbol"`
	TimeFrame string        `json:"timeframe"`
	Tick      []interface{} `json:"tick"` // time, bid, ask
}

type CurrentPrice struct {
	Symbol string
	Tick   Tick
}
