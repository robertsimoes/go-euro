package main

import (
	"fmt"

	"github.com/robertsimoes/go-euro/rates"
)

func main() {
	rates, _ := rates.AllCurrentRates()

	usd := rates.Rates["USD"]
	fmt.Printf("today 1 EUR = %s USD", usd)
}
