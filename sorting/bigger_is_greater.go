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

func quicksort(w *bufio.Writer, arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	N := len(arr)
	left := make([]int, 0, N)
	right := make([]int, 0, N)
	p := arr[0]
	for _, v := range arr {
		if v < p {
			left = append(left, v)
		}
		if v > p {
			right = append(right, v)
		}
	}
	left = quicksort(w, left)
	right = quicksort(w, right)
	left = append(left, p)
	left = append(left, right...)
	PrintArr(w, left)
	fmt.Fprintln(w, "")
	return left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	N := ReadInt(scanner)
	arr := ReadArr(scanner, N)
	quicksort(w, arr)
	w.Flush()
}
