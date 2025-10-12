package main

import "fmt"

func convertCurrency(from, to string, sum float64, operations *map[string]map[string]func(float64) float64) float64 {
	if from == to {
		return sum
	}

	if fromOps, ok := (*operations)[from]; ok {
		if converter, ok := fromOps[to]; ok {
			return converter(sum)
		}
	}

	fmt.Printf("Конвертация из %s в %s не найдена — возвращаем исходную сумму\n", from, to)
	return sum
}

func main() {
 	const usdEurRate = 1.07
	const usdRubRate = 94.0
	
	operations := map[string]map[string]func(float64) float64{
		"USD": {
			"EUR": func(sum float64) float64 { return sum / usdEurRate },
			"RUB": func(sum float64) float64 { return sum * usdRubRate },
		},
		"EUR": {
			"USD": func(sum float64) float64 { return sum * usdEurRate },
			"RUB": func(sum float64) float64 { return sum * usdEurRate * usdRubRate },
		},
		"RUB": {
			"USD": func(sum float64) float64 { return sum / usdRubRate },
			"EUR": func(sum float64) float64 { return sum / usdRubRate / usdEurRate },
		},
	}
	for {
	var currencyTo string
	var currencyFrom string
	var sum float64
	
	for {
		value, ok := getCurrencyFromToValue(currencyTo)
		if ok {
			currencyFrom = value
			break
		}
		fmt.Println("Попробуйте снова.")
	}

	for {
		value, ok := getCurrencyToValue()
		if ok {
			currencyTo = value;
			break 
		}
		fmt.Println("Попробуйте снова.")
	}


	for {
		value, ok := getSumToValue(currencyTo)
		if ok {
			sum = value
			break
		}
		fmt.Println("Попробуйте снова.")
	}

	value := convertCurrency(currencyFrom, currencyTo, sum, &operations)
		fmt.Printf("Вы конвертируете %d из %s в %s\n", sum, currencyFrom, currencyTo)
		fmt.Printf("Итоговый результат %f", value)
	break
}
}

func getSumToValue(targetCurrency string) (float64, bool) {
	var sum float64
	fmt.Printf("Какую сумму вы хотите перевести из %s: ", targetCurrency)
	fmt.Scan(&sum)
	if sum <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return 0, false
	}
	return sum, true
}

func getCurrencyToValue() (string, bool) {
	var currencyTo string
	fmt.Printf("В какую валюту Вы хотите конвертировать?: (USD/EUR/RUB) ")
	fmt.Scan(&currencyTo)

	if currencyTo != "USD" && currencyTo != "EUR" && currencyTo != "RUB" {
		fmt.Println("Значение невалидно. Возможные варианты: USD, RUB, EUR")
		return "", false
	}
	return currencyTo, true
}

func getCurrencyFromToValue(currencyTo string) (string, bool) {
	var currencyFrom string
	fmt.Println("Из какой валюты Вы хотите конвертировать?: (USD/EUR/RUB) ")
	fmt.Scan(&currencyFrom)
	if(currencyFrom != "USD" && currencyFrom != "EUR" && currencyFrom != "RUB"){
		fmt.Printf("Значение невалидно. Возможные варианты: USD, RUB, EUR")
	}
	if currencyFrom == currencyTo {
		fmt.Println("Валюты должны отличаться")
		return "", false
	}
	return currencyFrom, true
}
