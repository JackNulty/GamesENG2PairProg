// Author Shane, Jack & Darragh

package main

import (
	"fmt"
)

const count int = 8

// non decreasing, remove in place, return unique quantity
func removeDupes(numArray []int) (uniqueNums int) {
	var lastUnique = numArray[0]
	uniqueNums = 1 //first number is always unique

	var modifiedArray = [count]int{numArray[0], 0, 0, 0, 0, 0, 0, 0} //put first number into the array as it has to be unique
	var modArrayIndex = 1

	for index := 1; index < len(numArray); index++ {
		if numArray[index] != lastUnique {
			lastUnique = numArray[index]
			uniqueNums++
			modifiedArray[modArrayIndex] = numArray[index]
			modArrayIndex++
		}
	}
	fmt.Print(modifiedArray)
	return uniqueNums
}

func main() {
	arrayOfNums := [count]int{1, 2, 2, 4, 5, 5, 5, 6}
	uniques := removeDupes(arrayOfNums[:])
	fmt.Print("\nAmount of unique numbers: ")
	fmt.Print(uniques)
}
