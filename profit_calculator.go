package main

import "com.example/investement-calculator/finance"

//"fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)
	finance.Run()
	// ebt := finance.EarningBeforeTax(revenue, expenses)
	// profit := finance.NetProfit(ebt, taxRate)
	// ratio, _ := finance.EBTToProfitRatio(ebt, profit)

	// fmt.Printf("Earnings Before Tax: %.2f \nProfit: %.2f \nEBT To Profit Ratio: %.2f%%", ebt, profit, ratio)
}