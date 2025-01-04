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

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadPrices()
	if err != nil {
		return err
	}
	results := make(map[string]string)
	for _, price := range job.InputPrices {
		priceStr := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		results[fmt.Sprintf("%.2f", price)] = priceStr
	}
	return job.IOManager.WriteData(job)
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
