package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	fmt.Println("hi there")
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%v.json", taxRate*100))
		priceJob := prices.New(fm, taxRate)
		priceJob.Process()
	}

}
