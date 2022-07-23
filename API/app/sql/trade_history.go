package sql

import (
	"Appolo-api/app/config"
	"Appolo-api/app/models"
	"database/sql"
	"fmt"
)

func GetSumStrategyFromTradeHistory(c *config.Client, strategyName string) (models.TotalPNLAndCommission, error) {
	var dest models.TotalPNLAndCommission

	if err := c.DB.Get(&dest, `SELECT
	 SUM(pnl) AS pnl,
	 SUM(commission) AS commission 
	 FROM trade_history 
	 WHERE strategy_name = $1
	 `, strategyName); err != nil {
		if sql.ErrNoRows == err {
			return models.TotalPNLAndCommission{}, fmt.Errorf("no result on table trade_history for strategy_name %s : %w", strategyName, err)
		}
		return models.TotalPNLAndCommission{}, err
	}

	return dest, nil
}
