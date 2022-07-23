package models

type BinanceAccount struct {
	MakerCommission int  `json:"makerCommission"`
	TakerCommission int  `json:"takerCommission"`
	CanTrade        bool `json:"canTrade"`
	CanWithdraw     bool `json:"canWithdraw"`
	CanDeposit      bool `json:"canDeposit"`
	UpdateTime      int  `json:"updateTime"`
	Balances        []struct {
		Asset     string `json:"asset"`
		Available string `json:"free"`
		Locked    string `json:"locked"`
	} `json:"balances"`
	Permissions []string `json:"permissions"`
}
