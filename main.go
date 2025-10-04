package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)


var supportedOperations = [3]string{"avg", "sum", "med"}

func main(){

	for {
		operationType := getOperationType();

		if(operationType == ""){
			break;
		}
		numbers := getUserInput()
		calculatedResult := 0.0

		switch(operationType){
		case "avg": calculatedResult = getAverage(numbers)
		case "sum": calculatedResult = float64(getSum(numbers))
		case "med": calculatedResult = getMedian(numbers)
		default: calculatedResult = 0.0
		}

		fmt.Printf("Введенные значения: %v, вычисленное значение: %f\n", numbers, calculatedResult)

	}
}


func getUserInput()[]int {
	var input string
	fmt.Println("Введите числа через запятую (например: 1,2,3,4):")
	fmt.Scan(&input) 
	
	parts := strings.Split(input, ",")

	numbers := []int{}

		for _, part := range parts {
		part = strings.TrimSpace(part)
		if num, err := strconv.Atoi(part); err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}
func getMedian(numbers []int) float64 {
	n := len(numbers)
	if n == 0 {
		return 0
	}

 
	sorted := make([]int, n)
	copy(sorted, numbers)
 
	sort.Ints(sorted)

	if n%2 == 1 {
		return float64(sorted[n/2])
	}

	mid1 := sorted[n/2-1]
	mid2 := sorted[n/2]
	return float64(mid1+mid2) / 2.0
}

func getSum(numbers []int ) int{
	result := 0;

	for _, number := range numbers {
		result += number
	}

	return result
}

func getAverage(numbers []int) float64{
	result := 0;

	for _, number := range numbers {
		result += number
	}

	return float64(result) / float64(len(numbers))
}

func getOperationType() string {
	var operation string
	fmt.Print("Введите название операции: avg, sum или med")
	fmt.Scan(&operation)

	operation = strings.ToLower(strings.TrimSpace(operation))
	if operation == "" {
		return ""
	}

	for _, value := range supportedOperations {
		if operation == value {
			return operation
		}
	}

	fmt.Printf("Значение '%s' не поддерживается. Используйте: avg, sum, med", operation)
	return ""
}