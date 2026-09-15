package finance

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var filename string = "financial-info.txt"

type FromScanner struct {
	err error
}

func(fs *FromScanner) getInput(prompt string) (float64) {	
	if fs.err != nil {
		return 0
	}

	var value float64
	value, fs.err = GetUserInput(prompt)
	return value
}


func Run() () {
	sc := &FromScanner{}

	expenses := sc.getInput("Expenses: ")
	taxRate  := sc.getInput("Tax Rate: ")
	revenue  := sc.getInput("Revenue: ")
	
	if sc.err != nil {
		fmt.Printf("Input failed: %v\n", sc.err)
		return
	}
	
	ebt := EarningBeforeTax(revenue, expenses)
	profit := NetProfit(ebt, taxRate)
	ratio, _ := EBTToProfitRatio(ebt, profit)
	
	output := fmt.Sprintf("Earnings Before Tax: %.2f \nProfit: %.2f \nEBT To Profit Ratio: %.2f%%", ebt, profit, ratio)
	fmt.Printf("%s", output)
	
	WriteBalanceToFile(output, filename)
	
	text, _ := ReadBalanceFromFile(filename)
	fmt.Println(text)
}

func GetUserInput(prompt string) (float64, error) {
	fmt.Print(prompt)
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, fmt.Errorf("invalid input format: %w", err)
	}
	// value := strconv.ParseFloat(string(value), 64)
	if input <= 0 {
		message := fmt.Sprintf("%s cannot be <= 0, you entered %f",prompt, input)
		//panic(message)
		return 0, errors.New(message)
		
	}
	return input, nil
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