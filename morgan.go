package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type member struct {
	l int
	i int
	v string
}

func (m *member) end(i int) bool {
	return m.i+i >= m.l
}

func (m *member) next(i int) byte {
	if m.i+i >= m.l {
		return m.v[m.l-1]
	}
	return m.v[m.i+i]
}

func (m *member) pr_all() {
	for !m.end(0) {
		fmt.Printf("%c", m.next(0))
		m.advance(1)
	}
}

func (m *member) pr(i int) {
	for i > 0 && !m.end(0) {
		fmt.Printf("%c", m.next(0))
		m.advance(1)
		i--
	}
}

func (m *member) advance(k int) {
	m.i += k
}

type pair struct {
	s [2]member
}

func pr(x *member, y *member) {
	var c byte
	for {
		if x.end(0) || y.end(0) {
			break
		}
		a := x.next(0)
		b := y.next(0)
		if a < b {
			x.advance(1)
			c = a
		}
		if a > b {
			c = b
			y.advance(1)
		}
		if a == b {
			c = a
			x_k, y_k := pick_first(a, x, y)
			x.pr(x_k)
			y.pr(y_k)
			continue
		}
		fmt.Printf("%c", c)
	}
	x.pr_all()
	y.pr_all()
	fmt.Println("")
}

func pick_first(prev byte, x *member, y *member) (int, int) {
	_ = "breakpoint"
	k := 1
	for {
		if x.end(k) && y.end(k) {
			return k, k
		}
		if x.next(k) == y.next(k) {
			// XBCAA
			// XBCAB  BBBA
			//        BBBC
			// ZY
			// ZZ
			// CBABC AAAAAAX BAA
			// CBABX AAAAAAA BAX
			if prev < x.next(k) {
				return k, k
			}
			k++
			continue
		}
		x_k := 0
		y_k := 0
		min := x.next(k)
		if y.next(k) < min {
			min = y.next(k)
		}
		if prev < min {
			x_k = k
		}
		if prev < min {
			y_k = k
		}
		if x.next(k) < y.next(k) {
			return k, y_k
		} else {
			return x_k, k
		}
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	n := 0
	fmt.Scan(&n)
	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].s[0].v)
		arr[i].s[0].l = len(arr[i].s[0].v)
		fmt.Scan(&arr[i].s[1].v)
		arr[i].s[1].l = len(arr[i].s[1].v)
	}

	for i := 0; i < n; i++ {
		pr(&arr[i].s[0], &arr[i].s[1])
	}
}
