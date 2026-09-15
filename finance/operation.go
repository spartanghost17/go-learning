package finance

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Run() () {
	var revenue float64
	var expenses float64
	var taxRate float64

	var filename string = "financial-info.txt"

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
	WriteBalanceToFile(output, filename)
	text, _ := ReadBalanceFromFile(filename)
	fmt.Println(text)
}

func WriteBalanceToFile(balance string, filename string) error {
	return os.WriteFile(filename, []byte(balance), 0644)
}

func ReadBalanceFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(data)

	get := func(prefix string) string {
		for _, text := range strings.Split(text, "\n") {
			if strings.HasPrefix(text, prefix) {
				return strings.TrimSpace(strings.TrimPrefix(text, prefix))
			}
		}
		return ""
	}

	ebt := get("Earnings Before Tax:")
	profit := get("Profit:")
	ratio := get("EBT To Profit Ratio:")
	return fmt.Sprintf("Read from file, Earnings Before Tax: %s\nProfit: %s\nEBT To Profit Ratio: %s", ebt, profit, ratio), nil
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