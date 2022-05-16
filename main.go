package main

import (
	"fmt"
	"log"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/services"
)

func main() {
	client := services.GetBinanceClient()

	// check Binance API is up
	if err := services.Health(client); err != nil {
		log.Println(fmt.Sprintf("Binance API: %s", err.Error()))
	}
	log.Println("Binance API is working right!")

	
}
