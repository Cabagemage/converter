package main

import "fmt"

func UsdToEurConverter(usdCount float64, usdEurRate float64) float64 {
	return usdCount / usdEurRate
}

func UsdToRublesConverter(usdCount float64, usdRubRate float64) float64 {
	return usdCount * usdRubRate
}

func EuroToRublesConverter(euroCount float64, usdEurRate float64, usdRubRate float64) float64 {
	return euroCount * usdEurRate * usdRubRate
}

func main() {
 	const usdEurRate = 1.07
	const usdRubRate = 94.0

	euroFromUsd := UsdToEurConverter(100, usdEurRate)
	rublesFromUsd := UsdToRublesConverter(100, usdRubRate)
	rublesFromEur := EuroToRublesConverter(100, usdEurRate, usdRubRate)

	fmt.Printf("Euro from USD: %.2f\n", euroFromUsd)
	fmt.Printf("Rubles from USD: %.2f\n", rublesFromUsd)
	fmt.Printf("Rubles from EUR: %.2f\n", rublesFromEur)
}
