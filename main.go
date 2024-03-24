package main

import (
	"fmt"

	"github.com/cxhercules/go-practice-project/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(result)
}
