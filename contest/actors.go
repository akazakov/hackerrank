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

func count_p(s *[]int, p int) int {
	c := 0
	for _, v := range *s {
		if v < p {
			c++
		}
	}
	return c
}
func count_ps(s *[]int, p int) int {
	c := 0
	for i := 0; i < p; i++ {
		c += (*s)[i]
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	M := ReadInt(scanner)
	N := ReadInt(scanner)
	Q := ReadInt(scanner)
	segments := make([]int, M)
	ps := make([]int, N+1)
	actors := ReadArr(scanner, N)
	for _, v := range actors {
		segments[v]++
	}
	for _, v := range segments {
		ps[v]++
	}
	w = bufio.NewWriter(os.Stdout)
	for i := 0; i < Q; i++ {
		a := ReadInt(scanner)
		if a == 1 {
			n_i := ReadInt(scanner)
			m_i := ReadInt(scanner)
			prev := actors[n_i]
			actors[n_i] = m_i
			ps[segments[prev]]--
			ps[segments[m_i]]--
			segments[m_i]++
			segments[prev]--
			ps[segments[prev]]++
			ps[segments[m_i]]++
		} else {
			p := ReadInt(scanner)
			//fmt.Fprintln(w, count_p(&segments, p))
			fmt.Fprintln(w, count_ps(&ps, p))
		}
	}
	w.Flush()
}
