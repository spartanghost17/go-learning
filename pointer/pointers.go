package pointers

import "fmt"

func Run() {
	age := 10
	var agePtr *int = &age
	adultYears := GetAdultYears(agePtr)
	fmt.Println("Adult years:", adultYears)
	fmt.Println("Age:", age)
	fmt.Println("Age pointer:", agePtr)
	fmt.Println("Value pointed to:", *agePtr)
}

func GetAdultYears(age *int) (int) {
	*age = *age - 18;
	return *age
}
