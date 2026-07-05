package main

import (
	"bufio"
	"fmt"
	"os"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	_, err := fmt.Fscan(in, &n, &k)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &numbers[i])
	}

	fmt.Println(binarySearch(numbers, k))
}
