package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"taxRate"`
	InputPrices       []float64               `json:"inputPrices"`
	TaxIncludedPrices map[string]string       `json:"taxIncludedPrices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return
	}
	floatLines, err := conversion.StrigToFloat(lines)
	if err != nil {
		return
	}
	job.InputPrices = floatLines
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	results := make(map[string]string)
	for _, price := range job.InputPrices {
		priceStr := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		results[fmt.Sprintf("%.2f", price)] = priceStr
	}
	job.IOManager.WriteData(job)
}

func New(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: *fm,
		TaxRate:   taxRate,
	}
}
