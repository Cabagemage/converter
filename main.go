package main

import "fmt"

func convertCurrency(amount float64, from string, to string, usdEurRate float64, usdRubRate float64) float64 {
	switch from {
	case "USD":
		switch to {
		case "EUR":
			return amount / usdEurRate
		case "RUB":
			return amount * usdRubRate
		}
	case "EUR":
		switch to {
		case "USD":
			return amount * usdEurRate
		case "RUB":
			return amount * usdEurRate * usdRubRate
		}
	case "RUB":
		switch to {
		case "USD":
			return amount / usdRubRate
		case "EUR":
			return amount / usdRubRate / usdEurRate
		}
	}

	return amount
}

func main() {
 	const usdEurRate = 1.07
	const usdRubRate = 94.0
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

	value := convertCurrency(sum, currencyFrom, currencyTo, usdEurRate, usdRubRate)
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
