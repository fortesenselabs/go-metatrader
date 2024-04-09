package main

import (
	"fmt"

	"github.com/FortesenseLabs/go-metatrader/metatrader"
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

	// // get balance
	// balance, err := mt.GetBalance()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Print the response
	// fmt.Println(balance)

	// always get the current price
	// for {
	// 	// get current price
	// 	currentPrice, err := mt.GetCurrentPrice("Step Index")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// Print the response
	// 	fmt.Println(currentPrice)

	// 	time.Sleep(1 * time.Millisecond)
	// }

	// // get historical data
	// history, err := mt.GetHistoricalData("Step Index", timeframes.D1, actiontype.PRICE, "13-02-2021 00:00:00", "13-02-2024 00:00:00")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Print the response
	// fmt.Println(history)

	// get orders
	// orders, err := mt.GetOrders()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Print the response
	// fmt.Println(orders)

	// get positions
	// positions, err := mt.GetPositions()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Print the response
	// fmt.Println(positions)

	// currentPrice, err := mt.GetCurrentPrice("Step Index")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(currentPrice)

	// pips := 100
	// stopLevel := 0.1
	// normalizedPips := float64(pips) * stopLevel

	// // BUY
	// stopLoss := float64(currentPrice.Tick.Bid) - normalizedPips
	// takeProfit := float64(currentPrice.Tick.Bid) + normalizedPips

	// fmt.Println(stopLoss, takeProfit)

	// err = mt.Buy("Step Index", 0.1, stopLoss, takeProfit, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("BUY Order Sent")

}
