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
func partition(arr []int, lo, hi int) int {
	p := arr[hi]
	i := lo
	for j := lo; j <= hi-1; j++ {
		if arr[j] <= p {
			t := arr[i]
			arr[i] = arr[j]
			arr[j] = t
			i++
		}
	}
	t := arr[i]
	arr[i] = arr[hi]
	arr[hi] = t
	PrintArr(w, arr)
	fmt.Fprintln(w, "")
	return i
}

func quicksort(arr []int, lo, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		quicksort(arr, lo, p-1)
		quicksort(arr, p+1, hi)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	N := ReadInt(scanner)
	arr := ReadArr(scanner, N)
	quicksort(arr, 0, N-1)
	w.Flush()
}
