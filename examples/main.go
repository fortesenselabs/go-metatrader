package main

import (
	"fmt"

	"github.com/FortesenseLabs/go-metatrader/metatrader"
	"github.com/FortesenseLabs/go-metatrader/metatrader/actiontype"
	"github.com/FortesenseLabs/go-metatrader/metatrader/timeframes"
)

func main() {
	mt := metatrader.NewMetaTraderClient("localhost", 1122, false, false, "123456", []string{""})

	// Connect to the server
	mt.Connect()
	defer mt.Disconnect()

	// Send a request to the server

	// get account info
	accountInfo, err := mt.GetAccountInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response
	fmt.Println(accountInfo)

	// get balance
	balance, err := mt.GetBalance()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response
	fmt.Println(balance)

	// get historical data
	history, err := mt.GetHistoricalData("Step Index", timeframes.D1, actiontype.PRICE, "01-02-2024 00:00:00", "13-02-2024 00:00:00")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response
	fmt.Println(history)
}
