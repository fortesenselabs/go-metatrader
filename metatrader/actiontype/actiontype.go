package actiontype

const (
	PRICE  = "PRICE"
	TRADES = "TRADES" // Trades - only requires the actionType parameter
)

var ACTIONS = []string{
	"ORDER_TYPE_BUY",
	"ORDER_TYPE_SELL",
	"ORDER_TYPE_BUY_LIMIT",
	"ORDER_TYPE_SELL_LIMIT",
	"ORDER_TYPE_BUY_STOP",
	"ORDER_TYPE_SELL_STOP",
}
