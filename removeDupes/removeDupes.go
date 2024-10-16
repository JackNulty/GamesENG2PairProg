// Author Shane, Jack & Darragh

package main

import (
	"fmt"
)

// non decreasing, remove in place, return unique quantity
func removeDupes(numArray []int) (uniqueNums int) {
	var lastUnique = -999 //the most recent unique is set to be below -100 which is the minimum value stated to be valid
	uniqueNums = 0

	modifiedArray := make([]int, len(numArray))
	var modArrayIndex = 0

	for index := 0; index < len(numArray); index++ {
		if numArray[index] != lastUnique { //ensure the current number isnt the same as the most recent unique
			lastUnique = numArray[index] //update the most recent unique
			uniqueNums++
			modifiedArray[modArrayIndex] = numArray[index] //add current number to the new array without duplicates
			modArrayIndex++
		}
	}

	for index := 0; index < len(numArray); index++ { //modify array to so that the unique elements only are present and in order
		if index < uniqueNums {
			numArray[index] = modifiedArray[index] //repopulate the array to remove duplicates
		} else {
			numArray[index] = -999 //fill the rest of the array with filler
		}
	}

	return uniqueNums
}

func main() {
	arrayOfNums := [8]int{1, 1, 1, 1, 1, 1, 1, 1} //array is in order and has duplicates
	uniques := removeDupes(arrayOfNums[:])        //give back number of unique elements
	fmt.Print("\nAmount of unique numbers: ")
	fmt.Print(uniques)
}
