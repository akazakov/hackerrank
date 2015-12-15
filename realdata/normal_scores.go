package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func integrate(x float64) float64 {
	sum := x
	value := x
	for i := 1; i < 20; i++ {
		value = (value * x * x / (2.0*float64(i) + 1.0))
		sum = sum + value
	}
	result := 0.5 + (sum/math.Sqrt(2.0*math.Pi))*math.Exp(-(x*x)/2.0)
	return result
}

func convert(x float64) float64 {
	return x*200.0 + 2000.0
}

func f(x float64) float64 {
	return integrate((x-2000.0)/200.0) - 0.9
}

func newton(x0, x1, delta float64) float64 {
	f1 := f(x0)
	f2 := f(x1)
	if math.Abs(f1-f2) < delta {
		return x0
	}
	a := (x0 + x1) / 2.0
	if f(x0)*f(a) > 0 {
		return newton(a, x1, delta)
	} else {
		return newton(x0, a, delta)
	}
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(newton(2000.0*r.Float64(), 2000.0+2000.0*r.Float64(), 0.000001))
}
