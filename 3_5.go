package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// バブルソート、挿入ソート、選択ソートの実装

func bubbleSort(a []int, b []string, n int) ([]int, []string) {
	flag := 1

	for flag > 0 {
		flag = 0
		for i := n - 1; i > 0; i-- {
			if a[i - 1] > a[i] {
				tmp := a[i - 1]
				a[i - 1] = a[i]
				a[i] = tmp
				
				tmp_b := b[i - 1]
				b[i - 1] = b[i]
				b[i] = tmp_b

				flag = 1
			} 
		}
	}
	
	return a, b
}

func selectionSort(a []int, b []string, n int) ([]int, []string) {
	for i := 0; i < n; i++ {
		min_j := i

		for j := i; j < n; j++ {
			if a[j] < a[min_j] {
				min_j = j
			}
		}

		tmp := a[i]
		a[i] = a[min_j]
		a[min_j] = tmp

		tmp_b := b[i]
		b[i] = b[min_j]
		b[min_j] = tmp_b
	}
	
	return a, b
}

func insertSort(a []int, b []string, n int) ([]int, []string) {
	for i := 1; i < n; i++ {
		tmp_v := a[i]
		tmp_b_v := b[i]

		for j := i; j > 0; j-- {
			if a[j - 1] < tmp_v {
				a[j] = tmp_v
				b[j] = tmp_b_v
				break
			}
			a[j] = a[j - 1]
			b[j] = b[j - 1]
		}
	}
	
	return a, b
}

func returnInputValueSlice(scn *bufio.Scanner, n int) ([]int, []string) {
	var a []int
	var b []string

	for i := 0; i < n; i++ {
		scn.Scan()
		text := scn.Text()
		b = append(b, text)

		num, _ := strconv.Atoi(text[1:])
		a = append(a, num)
	}

	return a, b
}

func output(a []int, b[]string) {
	for i := range a {
		print(fmt.Sprintf("%v ", a[i]))
	}

	for i := range b {
		print(fmt.Sprintf("%v ", b[i]))
	}
}

func main() {
	var n int
	var a []int
	var b []string

	fmt.Scan(&n)
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	a, b = returnInputValueSlice(scn, n)

	a_bu, b_bu := bubbleSort(a, b, n)
	print("bubble sort\n")
	output(a_bu, b_bu)
	print("\n")

	a_se, b_se := selectionSort(a, b, n)
	print("selection sort\n")
	output(a_se, b_se)
	print("\n")

	a_in, b_in := insertSort(a, b, n)
	print("insert sort\n")
	output(a_in, b_in)
}