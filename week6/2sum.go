package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "sort"
import "math"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)
	numbersdict := make(map[int]bool, 0)
	for scanner.Scan() {
		lines := scanner.Text()
		number, err := strconv.Atoi(lines); checkError(err)
		numbers = append(numbers, number)
		numbersdict[number] = true
	}

	ts := make(map[int]bool, 0)
	sort.Ints(numbers)

	for ix := 0; ix < len(numbers); ix++ {
		x := numbers[ix]
		for iy := WhereIsY(numbers, x, 0, len(numbers));  iy < len(numbers); iy++ {
			y := numbers[iy]
			t := x + y

			if t > 10000 {
				break
			}

			if t >= -10000 && t<= 10000 && x != y {
				ts[t] = true
			}
		}
	}

	fmt.Println(len(ts))
}

func WhereIsY (array []int,x int, start int, end int) int {
	middle := start + int(math.Ceil(float64(end - start) / 2.0))
	if middle <= start || middle >= end {
		return start
	}

	y := array[middle]
	t := x + y

	//fmt.Println(x, y, t, middle)
	if t >= -10000 && t <= 10000 {
		return start
	} else if t < -10000 {
		//fmt.Println(middle, end)
		return WhereIsY(array, x, middle, end)
	} else {
		//fmt.Println(start, middle)
		return WhereIsY(array, x, start, middle)
	}
}

func checkError (err error) {
	if err != nil {
		panic(err)
	}
}
