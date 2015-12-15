package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

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

func processCase(scanner *bufio.Scanner) string {
	n := ReadInt(scanner)
	a := ReadArr(scanner, n)
	b := ReadArr(scanner, n)
	c := min(a[0], b[0])
	for i := range a {
		if min(a[i], b[i]) >= c {
			c = min(a[i], b[i])
		} else if max(a[i], b[i]) >= c {
			c = max(a[i], b[i])
		} else {
			return "NO"
		}
	}
	return "YES"
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := ReadInt(scanner)
	for i := 0; i < n; i++ {
		fmt.Println(processCase(scanner))
	}
}
