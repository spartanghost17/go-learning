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
	
	var hobies [3]string = [3]string{"Reading", "Traveling", "Cooking"}
	fmt.Println("Hobbies:", hobies)
	fmt.Println(hobies[0])
	fmt.Println(hobies[1:])
	
	slice1 := hobies[:2]
	slice2 := slice1[1:3]
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
	
	var goals []string = []string{"Learn Go", "Build a project"}
	fmt.Println("Goals:", goals)
	goals[1] = "Build a web app"
	fmt.Println("Updated Goals:", goals)
	goals = append(goals, "Contribute to open source")
	fmt.Println("Goals after append:", goals)

	var productsSlice []Product = []Product{
		{title: "Product A", id: "A1", price: 19.99},
		{title: "Product B", id: "B2", price: 29.99},
	}
	fmt.Println("Products Slice:", productsSlice)
	productsSlice = append(productsSlice, Product{title: "Product C", id: "C3", price: 39.99})

	var test1 string = "Hello, World!"
	fmt.Println("Test string:", test1[0:4])

}