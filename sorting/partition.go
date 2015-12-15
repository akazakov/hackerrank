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
		fmt.Fprintf(w, "%s%v", t, v)
		t = " "
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	N := ReadInt(scanner)
	arr := ReadArr(scanner, N)
	left := make([]int, 0, N)
	right := make([]int, 0, N)
	p := arr[0]
	for _, v := range arr {
		if v < p {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = append(left, right...)
	PrintArr(w, left)
	fmt.Fprintln(w, "")
	w.Flush()
}
