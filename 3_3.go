package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// バブルソートの実装

func bubbleSort(a []int, n int) []int {
	flag := 1
	var tmp int

	for flag > 0 {
		flag = 0
		for i := n - 1; i > 0; i-- {
			if a[i - 1] > a[i] {
				tmp = a[i]
				a[i] = a[i - 1]
				a[i - 1] = tmp

				flag = 1
			}
		}
	}

	return a
}

func returnInputValueSlice(scn *bufio.Scanner, n int) []int {
	var a []int

	for i := 0; i < n; i++ {
		scn.Scan()
		text := scn.Text()
		num, _ := strconv.Atoi(text)
		a = append(a, num)
	}

	return a
}

func output(a []int) {
	for i := range a {
		print(a[i])
	}
}

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	a = returnInputValueSlice(scn, n)
	a = bubbleSort(a, n)
	
	output(a)
}