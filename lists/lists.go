package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func Run() {
	var products [4]Product
	prices := [4]float64{19.99, 29.99, 39.99, 49.99}
	fmt.Println("Prices:", prices)
	fmt.Println("Prices length:", products)
	for i := range len(products) {
		products[i] = Product{
			title: fmt.Sprintf("Product %d", i+1),
			id:    fmt.Sprintf("P%d", i+1),
			price: prices[i],
		}
	}
	fmt.Println("Products:", products[0:2])
	var test [4]string = [4]string{"one", "two", "three", "four"}
	fmt.Println("Test:", test)

}