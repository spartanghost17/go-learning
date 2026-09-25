package lists

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

type transformFn func(int) int
type anotherFn func(int, []string, map[string][]int) ([]int, int)

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

	notes := make(map[string]string)
	notes["welcome"] = "Welcome to the notes app!"

	note, ok := notes["welcome"] // ok distinguishes a missing key from a zero value.
	fmt.Println("Note:", note, "Exists:", ok)

	delete(notes, "welcome")

	note1, ok1 := notes["welcome"] // ok1 distinguishes a missing key from a zero value.
	fmt.Println("Note:", note1, "Exists:", ok1)

	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{4, 5, 6, 7}

	dNumbers := transformNumbers(&numbers, double) //gets a callable
	tNumbers := transformNumbers(&numbers, triple) //gets a callable

	fmt.Println("Doubled numbers:", dNumbers)
	fmt.Println("Doubled numbers:", tNumbers)

	transformFn1 := getTransformerFunction(&numbers)     //returns a func
	transformFn2 := getTransformerFunction(&moreNumbers) //returns a func

	d1Numbers := transformNumbers(&numbers, transformFn1)     //gets a callable
	t1Numbers := transformNumbers(&moreNumbers, transformFn2) //gets a callable

	fmt.Println("\nDoubled numbers:", d1Numbers)
	fmt.Println("Doubled numbers:", t1Numbers)

	// anonymous function
	transformed4 := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println("Anonymous function:", transformed4)

	//---

	double := createTransformer(2)
	triple := createTransformer(3)

	fmt.Println("double:", transformNumbers(&numbers, double))
	fmt.Println("double:", transformNumbers(&numbers, triple))

	fact := factorial(3)
	fmt.Println("factorial:", fact)

	sum := sumup(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum:", sum, "\n")
	anoterSum := sumup(1, numbers...)
	fmt.Println("Another sum:", anoterSum)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func sumup(prefix int, numbers ...int) int {
	sum := 0
	fmt.Println("PrefixVal", prefix)
	for _, val := range numbers {
		sum += val
	}
	return sum
}
