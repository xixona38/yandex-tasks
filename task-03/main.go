package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var res = make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			arr[i], _ = strconv.Atoi(scanner.Text())
		}
	}

	for i := 1; i < n; i++ {
		if (arr[i] - arr[i-1]) != 2 {
			res = append(res, i)
		}
	}

	if len(res) == 1 {
		res = append(res, res[0])
	}

	for _, v := range res {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}
