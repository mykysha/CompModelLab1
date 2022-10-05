package main

import (
	"fmt"
	"math"

	"github.com/nnderhunov/CompModelLab1/math/rightangle"
	"github.com/nnderhunov/CompModelLab1/plotter"
)

func main() {
	a := 0.0
	b := math.Pi / 2
	step := math.Pi / 36
	eps := 0.0001
	f := func(x float64) float64 { return math.Log(math.Cos(x)) }

	var X, Y []float64

	stepsRequired := b / step

	for i := 0.0; i < stepsRequired-1; i++ {
		x := step * (i + 1)
		n := 2

		l := rightangle.RightRectangle(f, a, x, n)
		l2 := rightangle.RightRectangle(f, a, x, n*2)

		for math.Abs(l-l2)/3 > eps {
			n += 3
			l = rightangle.RightRectangle(f, 0, x, n)
			l2 = rightangle.RightRectangle(f, 0, x, n*2)
		}

		y := 0 - l2

		X = append(X, x)
		Y = append(Y, y)

		fmt.Printf("x = %dπ / 36,\ty = %f\n", int(i)+1, y)
	}

	p := plotter.NewPlotter("Lobachevsky function")

	err := p.Plot(X, Y, "-∫(0; x)ln(cos(t))dt [0; π/2]")
	if err != nil {
		panic(err)
	}

	err = p.Save("lobachevsky.png")
	if err != nil {
		panic(err)
	}
}
