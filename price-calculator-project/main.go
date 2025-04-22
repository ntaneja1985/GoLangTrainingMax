package main

import (
	"examples.com/price-calculator-project/filemanager"
	"fmt"

	//"examples.com/price-calculator-project/filemanager"
	"examples.com/price-calculator-project/prices"
	//"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process price job")
			fmt.Println(err)
		}
	}
}
