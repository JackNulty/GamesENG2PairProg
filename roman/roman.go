//Authors
// Mr Shane Trevor Gaming and Mr Jack P Sherman

package main

import (
	"fmt"
	"log"
)

func main() {
	var input string
	fmt.Print("Enter a Roman Numeral: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	number := generateNumbersFromString(input)
	fmt.Println("The integer value is:", number)
}

func generateNumbersFromString(Roman string) int {
	var answer int
	var numbers []int
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, char := range Roman {
		if value, exists := romanValues[char]; exists {
			numbers = append(numbers, value)
		}
	}
	for i := 0; i < len(numbers); i++ {
		if i < len(numbers)-1 && numbers[i] < numbers[i+1] {
			answer -= numbers[i]
		} else {
			answer += numbers[i]
		}
	}
	return answer
}
