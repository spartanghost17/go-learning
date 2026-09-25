package calculator

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"com.example/investement-calculator/calculator/filemanager"
)

// var filename string = "items_after_tax"
var taxFilename string = "taxe_prices.txt"

type IFileManager interface {
	WriteToFile(filename string, data []byte, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
}

type IItem interface {
	AfterTax() float64
}

type Item struct {
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	PriceAfterTax float64 `json:"price_after_tax"`
}

func (i *Item) AfterTax(taxRate float64) float64 {
	i.PriceAfterTax = i.Price + (i.Price * taxRate)
	return i.PriceAfterTax
}

func New(name string, price float64) *Item {
	return &Item{
		Name:  name,
		Price: price,
	}
}

func save(fileManager IFileManager, filename string, data []byte, perm os.FileMode) error {
	return fileManager.WriteToFile(filename, data, perm)
}

func read(fileManager IFileManager, filename string) ([]float64, error) {
	data, err := fileManager.ReadFile(filename)
	if err != nil {
		return []float64{}, fmt.Errorf("Error reading file %s: %w", filename, err)
	}
	text := string(data)
	tokens := strings.Split(text, ",")
	results := make([]float64, len(tokens))

	for i, val := range tokens {
		tokenFloat, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
		results[i] = tokenFloat
	}

	return results, nil
}

func Run() {
	fileManager := filemanager.New("v1")

	taxRates, err := read(fileManager, taxFilename) //[]float64{0, 10, 20}
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	items := []*Item{
		New("tomatoes", 10),
		New("Iphone", 20),
		New("Laptop", 30),
	}

	for _, taxRate := range taxRates {
		results := make([]*Item, len(items))

		for i, item := range items {
			result := New(item.Name, item.Price)
			result.AfterTax(taxRate / 100) // 10 means 10%, so use 0.10
			results[i] = result
		}

		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		outputFilename := fmt.Sprintf("items_after_tax_%g.json", taxRate)
		if err := save(fileManager, outputFilename, data, 0o644); err != nil {
			fmt.Println("Error saving file:", err)
			return
		}
	}
}
