package main

import (
	"testing"
	"fmt"
)

func TestSwap(t *testing.T) {
	array := []int {3, 2, 1, 0}

	Swap(array, 0, 3)
	if array[0] != 0 { t.FailNow() }
	if array[3] != 3 { t.FailNow() }

	Swap(array, 1, 2)
	if array[1] != 1 { t.FailNow() }
	if array[2] != 2 { t.FailNow() }
}

func TestPartitionArray(t *testing.T) {
	array := []int {3, 8, 2, 5, 1, 4, 7, 6}
	pivot := PartitionArray(array, 0)
	if pivot != 2 { t.FailNow() }

	array = []int {3, 8, 2, 5, 1, 4, 7, 6}
	pivot = PartitionArray(array, 5)
	if pivot != 3 { t.FailNow() }

	array = []int {3, 8, 2, 5, 1, 4, 7, 6}
	pivot = PartitionArray(array, 2)
	if pivot != 1 { t.FailNow() }
}

func TestQuickSortCountComparison(t *testing.T) {
	array := []int {3, 8, 2, 5, 1, 4, 7, 6}
	QuickSortCountComparison(array, func (arr []int) int { return 0 })

}

func TestIsMedian (t *testing.T) {
	if IsMedian(2, 0, 5) != true { t.FailNow() }
	if IsMedian(2, 3, 5) != false { t.FailNow() }
}
