package ichimoku

import (
	"Strategies/common"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func Worker() {
	log.Println("Starting.")

	client, err := common.NewClient()
	if err != nil {
		log.Println("Error generating credentials." + err.Error())
		return
	}

	data, err := client.GetHourlyDatas(2)
	if err != nil {
		log.Println("Error getting hourly datas." + err.Error())
		return
	}

	do := func() (common.FuturesPosition, string) {
		fp, err := client.GetFuturesPosition("test_ichimoku")
		if err != nil {
			log.Println("Error getting futures position." + err.Error())
			return common.FuturesPosition{}, ""
		}
		quantity, err := client.Convert("USDT", "BTC", "2000")
		if err != nil {
			log.Println("Error converting quantity." + err.Error())
			return common.FuturesPosition{}, ""
		}
		qty := fmt.Sprintf("%.3f", quantity)

		return fp, qty
	}

	// bullish cross & bearish cross from baseline to conversion line
	if data[1].Ichimoku[0] <= data[1].Ichimoku[1] && data[0].Ichimoku[0] > data[0].Ichimoku[1] {
		fp, qty := do()
		if fp.EntryPrice == "0.0" {
			if err := client.BuyAndSell("sell", "ichimoku", qty); err != nil {
				log.Println("Error buying or selling." + err.Error())
				return
			}
			fp.Quantity = qty
		}

		// Convert quantities to string to add them
		a, err := strconv.ParseFloat(fp.Quantity, 64)
		if err != nil {
			log.Println("Error converting quantity." + err.Error())
			return
		}
		b, err := strconv.ParseFloat(qty, 64)
		if err != nil {
			log.Println("Error converting quantity." + err.Error())
			return
		}

		if err := client.BuyAndSell("buy", "ichimoku", fmt.Sprintf("%f", a+b)); err != nil {
			log.Println("Error buying or selling." + err.Error())
			return
		}
	} else if data[1].Ichimoku[0] >= data[1].Ichimoku[1] && data[0].Ichimoku[0] < data[0].Ichimoku[1] {
		fp, qty := do()
		if fp.EntryPrice == "0.0" {
			if err := client.BuyAndSell("buy", "ichimoku", qty); err != nil {
				log.Println("Error buying or selling." + err.Error())
				return
			}
			fp.Quantity = qty
		}

		// Convert quantities to string to add them
		a, err := strconv.ParseFloat(fp.Quantity, 64)
		if err != nil {
			log.Println("Error converting quantity." + err.Error())
			return
		}
		b, err := strconv.ParseFloat(qty, 64)
		if err != nil {
			log.Println("Error converting quantity." + err.Error())
			return
		}

		if err := client.BuyAndSell("sell", "ichimoku", fmt.Sprintf("%f", a+b)); err != nil {
			log.Println("Error buying or selling." + err.Error())
			return
		}
	}

	log.Println("Done.")
}
