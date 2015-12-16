package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var w *bufio.Writer

func ReadInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func ReadArr(scanner *bufio.Scanner, n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = ReadInt(scanner)
	}
	return result
}

func PrintArr(w *bufio.Writer, arr []int) {
	t := ""
	for _, v := range arr {
		fmt.Fprint(w, t)
		fmt.Fprint(w, v)
		t = " "
	}
}
func partition(arr []int, lo, hi int) (int, int) {
	count := 0
	p := arr[hi]
	i := lo
	for j := lo; j <= hi-1; j++ {
		if arr[j] <= p {
			t := arr[i]
			arr[i] = arr[j]
			arr[j] = t
			i++
			count++
		}
	}
	t := arr[i]
	arr[i] = arr[hi]
	arr[hi] = t
	count++
	return i, count
}
func insertion_sort(arr []int) int {
	N := len(arr)
	count := 0
	for j := 1; j < N; j++ {
		i := j
		c := arr[i]
		for ; i > 0; i-- {
			if c < arr[i-1] {
				arr[i] = arr[i-1]
				count++
			} else {
				break
			}
		}
		arr[i] = c
	}
	return count
}

func quicksort(arr []int, lo, hi int) int {
	count := 0
	if lo < hi {
		p, c := partition(arr, lo, hi)
		count += c
		count += quicksort(arr, lo, p-1)
		count += quicksort(arr, p+1, hi)
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	N := ReadInt(scanner)
	arr := ReadArr(scanner, N)
	arr_copy := make([]int, N)
	copy(arr_copy, arr)

	i := insertion_sort(arr_copy)
	q := quicksort(arr, 0, N-1)
	fmt.Fprintln(w, i-q)

	w.Flush()
}
