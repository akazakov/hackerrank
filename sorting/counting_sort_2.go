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
	arr := ReadArr(scanner, N)
	counts := make([]int, 100)
	for _, v := range arr {
		counts[v]++
	}
	for k, v := range counts {
		for i := v; i > 0; i-- {
			fmt.Fprint(w, k)
			fmt.Fprint(w, " ")
		}
	}
	w.Flush()
}
