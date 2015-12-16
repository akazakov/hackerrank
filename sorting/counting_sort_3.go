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
	for i := 0; i < N; i++ {
		arr[i] = ReadInt(scanner)
		scanner.Scan()
		_ = scanner.Text()
	}
	counts := make([]int, 100)
	for _, v := range arr {
		counts[v]++
	}
	p := 0
	for _, v := range counts {
		p += v
		fmt.Fprint(w, p)
		fmt.Fprint(w, " ")
	}
	w.Flush()
}
