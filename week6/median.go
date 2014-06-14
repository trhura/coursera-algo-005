package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "container/heap"

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)
	for scanner.Scan() {
		lines := scanner.Text()
		number, err := strconv.Atoi(lines); checkError(err)
		numbers = append(numbers, number)
	}
	fmt.Println("Median = ", HeapCountMedian(numbers))
}

func HeapCountMedian (numbers []int) int {
	maxheap := &MaxHeap{}
	minheap := &MinHeap{}
	heap.Init(maxheap)
	heap.Init(minheap)

	sum_of_median := 0
	num1 := numbers[0]
	num2 := numbers[1]
	sum_of_median += num1

	if (num1 < num2) {
		heap.Push(maxheap, num1)
		heap.Push(minheap, num2)
		sum_of_median += num1
	} else {
		heap.Push(maxheap, num2)
		heap.Push(minheap, num1)
		sum_of_median += num2
	}

	for _, num := range numbers[2:] {
		if (num < []int(*maxheap)[0]) {
			heap.Push(maxheap, num)
		} else {
			heap.Push(minheap, num)
		}

		if (minheap.Len() + 1 < maxheap.Len()) {
			heap.Push(minheap, heap.Remove(maxheap, 0))
		}

		if (maxheap.Len() + 1 < minheap.Len()) {
			heap.Push(maxheap, heap.Remove(minheap, 0))
		}

		median := 0
		if (minheap.Len() > maxheap.Len()) {
			median = []int(*minheap)[0]
		} else {
			median = []int(*maxheap)[0]
		}
		sum_of_median += median
	}

	return sum_of_median
}

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func checkError (err error) {
	if err != nil {
		panic(err)
	}
}
