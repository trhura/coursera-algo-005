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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	array := make([]int, 0)

	/* Scan each line, parse number  and append it to array */
	for scanner.Scan() {
		string := scanner.Text()
		integer, err := strconv.Atoi(string); Check(err)
		array = append(array, integer)
	}

	_, count := CountInversions(array)
	fmt.Println(count)
}

func CountInversions(array []int) ([]int, int) {
	count  := 0
	sorted := make([]int, len(array))
	size   := len(array)

	if size <= 1 {
		return array, count
	}

	half := size / 2
	fst_half := array[0:half]
	snd_half := array[half:size]

	sorted_fst_half, fst_count := CountInversions(fst_half)
	sorted_snd_half, snd_count := CountInversions(snd_half)
	count += fst_count + snd_count

	j := 0
	k := 0
	for i, _ := range sorted {
		if (j >= len(sorted_fst_half)) {
			sorted[i] = sorted_snd_half[k]
			k++
			continue
		}

		if (k >= len(sorted_snd_half)) {
			sorted[i] = sorted_fst_half[j]
			j++
			continue
		}

		if sorted_snd_half[k] < sorted_fst_half[j] {
			sorted[i] = sorted_snd_half[k]
			k++

			count += len(sorted_fst_half) - j
			continue
		}

		sorted[i] = sorted_fst_half[j]
		j++
	}

	return sorted, count
}
