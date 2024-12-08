package main

import "fmt"

func main() {
	age := 24
	//integer point
	agePointer := &age
	fmt.Println("Age: ", agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Years", adultYears, age)
	editAgeToAdultYears(agePointer)
	fmt.Println("Edited Age: ", age)
}

func getAdultYears(age *int) int {
	//pointer as value
	return *age - 18
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
