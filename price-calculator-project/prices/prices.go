package prices

import (
	"examples.com/price-calculator-project/conversion"
	"examples.com/price-calculator-project/iomanager"
	"fmt"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) Process() error {
	//First load the data inside Input Prices
	//Note we are passing pointers so that we work on the original job not on its copy
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", taxIncludedPrice), 64)
	}
	//fmt.Println(result)
	//return result
	//Write to a file
	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

// NewTaxIncludedPriceJob: Constructor Function
// NewTaxIncludedPriceJob: In Go, when you access a struct field through a pointer, Go automatically dereferences the pointer for you
func NewTaxIncludedPriceJob(taxRate float64, fm iomanager.IOManager) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: fm,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	//fmt.Println(job.InputPrices)
	job.InputPrices = prices
	return nil
}
