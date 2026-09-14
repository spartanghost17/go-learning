package finance

import "errors"

func EarningBeforeTax(revenue, expenses float64) (float64) {
	return revenue - expenses
}

func NetProfit(ebt, taxRate float64) (float64) {
	return ebt - (1 - taxRate/100)
}

func EBTToProfitRatio(ebt, profit float64) (float64, error) {
	if profit <= 0.0 {
		return 0, errors.New("Cannot calculate ratio with profit <= 0")
	}
	return (ebt / profit) * 100, nil
}