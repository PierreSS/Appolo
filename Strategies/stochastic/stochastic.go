package stochastic

import (
	"Strategies/common"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Worker() {
	log.Println("Starting.")

	client, err := common.NewClient()
	if err != nil {
		log.Println("Error generating credentials." + err.Error())
		return
	}

	data, err := client.GetHourlyDatas(3)
	if err != nil {
		log.Println("Error getting hourly datas." + err.Error())
		return
	}

	fp, err := client.GetFuturesPosition("test_stochastic")
	if err != nil {
		log.Println("Error getting futures position." + err.Error())
		return
	}

	do := func() string {
		quantity, err := client.Convert("USDT", "BTC", "2000")
		if err != nil {
			log.Println("Error converting quantity." + err.Error())
			return ""
		}
		qty := fmt.Sprintf("%.3f", quantity)
		return qty
	}

	if fp.EntryPrice == "0.0" {
		if data[0].Stochastic > 55 && (data[1].Stochastic < 50 || data[2].Stochastic < 50) {
			if err := client.BuyAndSell("buy", "stochastic", do()); err != nil {
				log.Println("Error buying or selling." + err.Error())
				return
			}
		} else if data[0].Stochastic < 45 && (data[1].Stochastic > 50 || data[2].Stochastic > 50) {
			if err := client.BuyAndSell("sell", "stochastic", do()); err != nil {
				log.Println("Error buying or selling." + err.Error())
				return
			}
		}
	}

	if fp.EntryPrice != "0.0" {
		if data[1].Supertrend == "short" && data[0].Supertrend == "long" {
			if err := client.BuyAndSell("buy", "stochastic", fp.Quantity); err != nil {
				log.Println("Error buying or selling." + err.Error())
				return
			}
		}
		if data[1].Supertrend == "long" && data[0].Supertrend == "short" {
			if err := client.BuyAndSell("sell", "stochastic", fp.Quantity); err != nil {
				log.Println("Error buying or selling." + err.Error())
				return
			}
		}
	}

	log.Println("Done.")
}
