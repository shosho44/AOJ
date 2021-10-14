package main

import "fmt"

// 挿入ソートの実装

func printSlice(a []int) {
	for _, a_v := range a {
		print(a_v)
	}
}

func insersionSort(a []int, n int) {
	for i := 1; i < n; i++ {
		j := i - 1
		sl_content := a[i]
		
		for j >= 0 {
			if a[j] <= sl_content {
				break
			}
			
			a[j+1] = a[j]
			j -= 1
		}

		a[j+1] = sl_content
	}

	printSlice(a)
}

func main() {
	var inp, n int
	var a []int
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&inp)
		a = append(a, inp)
	}

	insersionSort(a, n)
}