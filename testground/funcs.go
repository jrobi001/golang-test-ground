package main

import (
	"fmt"
	"math"
	"time"
)

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func euclidian(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func sumDistance2D(dist func(x, y float64) float64, points [][]float64) (sum float64, ok bool) {

	if len(points[0]) != len(points[1]) {
		sum -= 1
		return sum, ok
	}
	ok = true

	for i, _ := range points[0] {
		sum += dist(points[0][i], points[1][i])
	}

	return sum, ok
}

func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return (math.Pi * math.Pow(c.radius, 2))
}

func info(s shape) {
	fmt.Println(s.area())
}

func durationMicroseconds() func() string {
	start := time.Now()
	return func() string {
		t := time.Since(start)
		elapsed := fmt.Sprintln("Microseconds:", t.Microseconds())
		return elapsed
	}
}

func main() {
	// x := []float64{2, 4, 5, 1, 2, 32}
	// y := []float64{6, 23, 1, 52, 5, 23}

	// z := [][]float64{x, y}

	// sumZ, ok := sumDistance2D(euclidian, z)

	// fmt.Println(sumZ, ok)

	c := circle{
		radius: 4.0,
	}
	s := square{
		length: 4.0,
		width:  2.0,
	}

	fmt.Println(c.area(), s.area())
	info(c)
	info(s)

	t := durationMicroseconds()
	fmt.Print(t())
	fmt.Print(t())
	fmt.Print(t())
	fmt.Print(t())

	a := 42
	b := &a
	a++
	fmt.Println(a, *b)
	*b++
	fmt.Println(a, *b)

	func incrementInt(a *int) {
		*a++
	}

	g := 42
	incrementInt(&g)
	fmt.Println(g)

}



func sumInts(a ...int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
