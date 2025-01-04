package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"taxRate"`
	InputPrices       []float64           `json:"inputPrices"`
	TaxIncludedPrices map[string]string   `json:"taxIncludedPrices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	floatLines, err := conversion.StrigToFloat(lines)
	if err != nil {
		return err
	}
	job.InputPrices = floatLines
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadPrices()
	if err != nil {
		// return err
		errChan <- err
		close(errChan)
		return
	}
	results := make(map[string]string)
	for _, price := range job.InputPrices {
		priceStr := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		results[fmt.Sprintf("%.2f", price)] = priceStr
	}
	fmt.Println(results)
	job.IOManager.WriteData(job)
	doneChan <- true
	close(doneChan)
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
