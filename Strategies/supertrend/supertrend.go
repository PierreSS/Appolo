package supertrend

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
		fp, err := client.GetFuturesPosition("test_supertrend")
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

	if data[1].Supertrend == "long" && data[0].Supertrend == "short" {
		fp, qty := do()
		if fp.EntryPrice == "0.0" {
			if err := client.BuyAndSell("buy", "supertrend", qty); err != nil {
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

		if err := client.BuyAndSell("sell", "supertrend", fmt.Sprintf("%f", a+b)); err != nil {
			log.Println("Error buying or selling." + err.Error())
			return
		}
	} else if data[1].Supertrend == "short" && data[0].Supertrend == "long" {
		fp, qty := do()
		if fp.EntryPrice == "0.0" {
			if err := client.BuyAndSell("sell", "supertrend", qty); err != nil {
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

		if err := client.BuyAndSell("buy", "supertrend", fmt.Sprintf("%f", a+b)); err != nil {
			log.Println("Error buying or selling." + err.Error())
			return
		}
	}

	log.Println("Done.")
}
