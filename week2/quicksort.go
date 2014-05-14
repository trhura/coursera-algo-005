package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func Check(err error) {
	if (err != nil) {
		panic(err)
	}
}

func main () {
	array := ScanArrayFromStdin()
	count := 0

	array1 := make([]int, len(array))
	copy(array1, array)
	count = QuickSortCountComparison(array1, func (arr []int) int {
		// use first element
		return 0
	})
	fmt.Println("Pivot FST Comparison Count: ", count)

	array2 := make([]int, len(array))
	copy(array2, array)
	count = QuickSortCountComparison(array2, func (arr []int) int {
		// use last element
		return len(arr) - 1
	})
	fmt.Println("Pivot LST Comparison Count: ", count)

	array3 := make([]int, len(array))
	copy(array3, array)
	count = QuickSortCountComparison(array3, func (arr []int) int {
		// use the median of first, last and middle elements
		fst := arr[0]
		lst := arr[len(arr)-1]
		mdl := arr[(len(arr)-1)/2]

		if IsMedian(fst, mdl, lst) {
			return 0
		}

		if IsMedian(lst, fst, mdl) {
			return len(arr) - 1
		}

		if IsMedian(mdl, fst, lst) {
			return (len(arr) - 1) / 2
		}

		panic("It shouldn't reach here.")
	})
	fmt.Println("Pivot MDN Comparison Count: ", count)
}

func ScanArrayFromStdin () []int {
	scanner := bufio.NewScanner(os.Stdin)
	array := make([]int, 0)

	/* Scan each line, parse number  and append it to array */
	for scanner.Scan() {
		string := scanner.Text()
		integer, err := strconv.Atoi(string); Check(err)
		array = append(array, integer)
	}
	return array
}

type PivotCallback func(array []int) (pivot int)

func QuickSortCountComparison (array []int, pivotPicker PivotCallback) int {
	count := 0
	if len(array) <= 1 {
		return count
	}

	pivot := pivotPicker(array)
	pivot  = PartitionArray(array, pivot)
	count = len(array) - 1

	lhsArray := array[:pivot]
	rhsArray := array[pivot+1:]
	count += QuickSortCountComparison(lhsArray, pivotPicker)
	count += QuickSortCountComparison(rhsArray, pivotPicker)

	return count
}

func PartitionArray (array []int, pivot int) int {
	Swap(array, pivot, 0)
	pivot = 1

	for j := 1; j < len(array); j++ {
		if array[j] < array[0] {
			Swap(array, pivot, j)
			pivot++
		}
	}
	pivot--

	Swap(array, 0, pivot)
	return pivot
}

func Swap (array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func IsMedian (m int, j int , k int) bool {
	isBetween := func (m int, lower int, upper int) bool {
		return m <= upper && m >= lower
	}

	return isBetween(m, j, k) || isBetween(m, k, j)
}
