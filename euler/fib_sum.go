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

func FindSum(n int, base int) int {
	ceil := (n - 1) / base
	sum := base * (ceil * (ceil + 1)) / 2
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	T := ReadInt(scanner)
	for i := 0; i < T; i++ {
		n := ReadInt(scanner)
		fmt.Fprintln(w, FindSum(n, 3)+FindSum(n, 5)-FindSum(n, 15))
	}
	w.Flush()
}
