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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	N := ReadInt(scanner)
	arr := make([]int, N)
	strings := make([]string, N)
	result := make([]string, N)
	for i := 0; i < N; i++ {
		arr[i] = ReadInt(scanner)
		scanner.Scan()
		strings[i] = scanner.Text()
	}
	counts := make([]int, 100)
	starts := make([]int, 100)
	for _, v := range arr {
		counts[v]++
	}
	p := 0
	for i, v := range counts {
		starts[i] = p
		p += v
	}
	t := len(arr) / 2
	for i, s := range strings {
		k := arr[i]
		start := starts[k]
		if i >= t {
			result[start] = s
		} else {
			result[start] = "-"
		}
		starts[k]++
	}
	for _, s := range result {
		fmt.Fprint(w, s)
		fmt.Fprint(w, " ")
	}

	w.Flush()
}
