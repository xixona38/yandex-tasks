package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isPrime(n int) bool {
	x := int(math.Sqrt(float64(n)))
	for j := 2; j <= x; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	simple := make(map[int]bool)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			simple[i] = true
		}
	}

	if ok := simple[n]; ok {
		fmt.Println("1")
		fmt.Println(n)
		return
	}

	for k := range simple {
		other := n - k
		if ok := simple[other]; ok {
			fmt.Println(2)
			fmt.Println(k, other)
			return
		}
	}

	q := n - 3
	for k := range simple {
		other := q - k
		if ok := simple[other]; ok {
			fmt.Println(3)
			fmt.Println(3, k, other)
			return
		}
	}
}
