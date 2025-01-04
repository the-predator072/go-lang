package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))
	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%v.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.New(fm, taxRate)
		go priceJob.Process(doneChans[i], errChans[i])
	}

	for i, _ := range taxRates {
		select {
		case err := <-errChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}
	// for _, errChan := range errChans {
	// 	<-errChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
