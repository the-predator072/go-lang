package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Couldn't read the file")
		fmt.Println(err)
		file.Close()
		return
	}
	scanner := bufio.NewScanner(file)
	lines := []float64{}
	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Failed to parse the text")
			fmt.Println(err)
			file.Close()
			return
		}
		lines = append(lines, floatPrice)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Couldn't read the file")
		fmt.Println(err)
		file.Close()
		return
	}
	job.InputPrices = lines
}

func (job *TaxIncludedPriceJob) Process() {
	results := make(map[string]float64)

	for _, price := range job.InputPrices {
		results[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(results)
}

func New(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
