package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var w *bufio.Writer

func ReadInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func ReadWord(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
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

func FindBigger(s string) bool {
	l := len(s)
	for i := l - 1; i >= 1; i-- {
		if s[i-1] < s[i] {
			p := s[i-1]
			k := i
			for j := i + 1; j < l; j++ {
				if s[j] > p && s[j] <= s[k] {
					k = j
				}
			}
			result := s[:i-1] + s[k:k+1]
			ps := strings.Split(s[i:], "")
			ps[k-i] = s[i-1 : i]
			sort.Strings(ps)
			result += strings.Join(ps, "")
			fmt.Fprintln(w, result)
			return true
		}
	}
	return false
}

func PrintStr(s string, j, k int) {
	jj := s[j]
	kk := s[k]
	for i := 0; i < len(s); i++ {
		if i == j {
			fmt.Fprintf(w, "%c", kk)
		} else if i == k {
			fmt.Fprintf(w, "%c", jj)
		} else {
			fmt.Fprintf(w, "%c", s[i])
		}
	}
	fmt.Fprintln(w, "")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w = bufio.NewWriter(os.Stdout)
	T := ReadInt(scanner)
	for i := 0; i < T; i++ {
		s := ReadWord(scanner)
		if !FindBigger(s) {
			fmt.Fprintln(w, "no answer")
		}
	}
	w.Flush()
}
