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

func PrintArr(w *bufio.Writer, arr *[]int) {
	t := ""
	for _, v := range *arr {
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
	c := arr[N-1]
	i := N - 1
	for ; i > 0; i-- {
		if c < arr[i-1] {
			arr[i] = arr[i-1]
			PrintArr(w, &arr)
			fmt.Fprintln(w, "")
		} else {
			break
		}
	}
	arr[i] = c
	PrintArr(w, &arr)
	w.Flush()
}
