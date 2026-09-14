package main

import (
	"fmt"

	"com.example/investement-calculator/finance"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	
	ebt := finance.EarningBeforeTax(revenue, expenses)
	profit := finance.NetProfit(ebt, taxRate)
	ratio, _ := finance.EBTToProfitRatio(ebt, profit)

	fmt.Printf("Earnings Before Tax: %.2f \nProfit: %.2f \nEBT To Profit Ratio: %.2f%%", ebt, profit, ratio)
}