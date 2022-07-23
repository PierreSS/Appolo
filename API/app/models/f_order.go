package models

type FuturesOrder struct {
	Symbol      string `json:"symbol"`
	OrderID     int64  `json:"orderId" swaggertype:"number"`
	UpdateTime  int64  `json:"updateTime"`
	OrigQty     string `json:"origQty"`
	Price       string `json:"price"`
	Type        string `json:"type"`
	Side        string `json:"side"`
	Commission  string
	RealizedPNL string
}
