package atclosing

import (
	_ "github.com/lib/pq"
)

// func Worker() {
// 	log.Println("Starting.")

// 	client, err := common.NewClient()
// 	if err != nil {
// 		log.Println("Error generating credentials." + err.Error())
// 		return
// 	}

// 	data, err := client.GetHourlyData()
// 	if err != nil {
// 		log.Println("Error getting hourly datas." + err.Error())
// 		return
// 	}

// 	if data.CloseTime.Hour() == 23 {
// 		quantity, err := client.Convert("USDT", "BTC", "2000")
// 		if err != nil {
// 			log.Println("Error converting quantity." + err.Error())
// 			return
// 		}

// 		f, err := common.ReadConfig()
// 		if err != nil {
// 			log.Println("Error reading quantity from config.json." + err.Error())
// 			return
// 		}
// 		f.AtClosingBoughtQuantity = fmt.Sprintf("%.3f", quantity)
// 		if err := common.WriteConfig(f); err != nil {
// 			log.Println("Error writing quantity in config.json." + err.Error())
// 			return
// 		}

// 		if err := client.BuyAndSell("buy", "atclosing", f.AtClosingBoughtQuantity); err != nil {
// 			log.Println("Error buying or selling." + err.Error())
// 			return
// 		}
// 	}
// 	if data.CloseTime.Hour() == 0 {
// 		f, err := common.ReadConfig()
// 		if err != nil {
// 			log.Println("Error reading quantity from config.json." + err.Error())
// 			return
// 		}

// 		if err := client.BuyAndSell("sell", "atclosing", f.AtClosingBoughtQuantity); err != nil {
// 			log.Println("Error buying or selling." + err.Error())
// 			return
// 		}
// 	}

// 	log.Println("Done.")
// }
