package models

// //BigInt big.Int alias
// type BigInt struct {
// 	big.Int
// }

// // Value implements the Valuer interface for BigInt
// func (b *BigInt) Value() (driver.Value, error) {
// 	if b != nil {
// 		return b.String(), nil
// 	}
// 	return nil, nil
// }

// // Scan implements the Scanner interface for BigInt
// func (b *BigInt) Scan(value interface{}) error {
// 	var i sql.NullString
// 	if err := i.Scan(value); err != nil {
// 		return err
// 	}
// 	if _, ok := b.SetString(i.String, 10); ok {
// 		return nil
// 	}
// 	return fmt.Errorf("Could not scan type %T into BigInt", value)
// }

// Order define an order buy or sell response from binance
type Order struct {
	Symbol          string `json:"symbol"`
	OrderID         string `json:"orderId" swaggertype:"number"`
	TransactionTime int    `json:"transactTime"`
	BoughtQuantity  string `json:"executedQty"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Side            string `json:"side"`
	Fills           []struct {
		Price      string `json:"price"`
		Quantity   string `json:"qty"`
		Commission string `json:"commission"`
	} `json:"fills"`
}
