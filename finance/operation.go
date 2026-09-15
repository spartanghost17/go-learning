package finance

import (
	"errors"
	"fmt"
	"os"
)

func Run() () {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	
	ebt := EarningBeforeTax(revenue, expenses)
	profit := NetProfit(ebt, taxRate)
	ratio, _ := EBTToProfitRatio(ebt, profit)
	output := fmt.Sprintf("Earnings Before Tax: %.2f \nProfit: %.2f \nEBT To Profit Ratio: %.2f%%", ebt, profit, ratio)
	fmt.Printf("%s", output)
	WriteBalanceToFile(output, "financial-info.txt")
}

func WriteBalanceToFile(balance string, filename string) error {
	return os.WriteFile(filename, []byte(balance), 0644)
}

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