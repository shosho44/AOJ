package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 選択ソートの実装

func selectionSort(a []int, n int) []int {
	for i := 0; i < n; i++ {
		var tmp int

		min_j := i
		for j := i; j < n; j++ {
			if a[min_j] < a[j] {
				min_j = j
			}
		}
		tmp = a[i]
		a[i] = a[min_j]
		a[min_j] = tmp
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
		print(fmt.Sprintf("%v ", a[i]))
	}
}

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	a = returnInputValueSlice(scn, n)
	a = selectionSort(a, n)
	
	output(a)
}