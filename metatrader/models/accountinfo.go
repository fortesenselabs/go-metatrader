package models

type AccountInfo struct {
	Account        string  `json:"account"`
	Balance        float64 `json:"balance"`
	Error          bool    `json:"error"`
	Broker         string  `json:"broker"`
	Currency       string  `json:"currency"`
	Server         string  `json:"server"`
	TradingAllowed int64   `json:"trading_allowed"`
	BotTrading     int64   `json:"bot_trading"`
	Equity         float64 `json:"equity"`
	Margin         float64 `json:"margin"`
	MarginFree     float64 `json:"margin_free"`
	MarginLevel    float64 `json:"margin_level"`
	Time           string  `json:"time"`
}
