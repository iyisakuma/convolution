package tp

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
	"math/rand"
	"sync"
)

const (
	a         = 0.8
	maxPoints = 200
)

func s(n float64) float64 { return math.Cos(n*0.08) + 0.3*math.Sin(0.2*n) }
func r() float64          { return (rand.Float64() * 0.6) - 0.3 }
func h(n float64) float64 {
	if n < 0 {
		return 0
	}
	return (1 - a) * math.Pow(a, float64(n))
}
func convolution1D(x, h func(float64) float64) [maxPoints]float64 {
	convolution1DResults := [maxPoints]float64{}

	wg := sync.WaitGroup{}
	for i := range convolution1DResults {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			var temp float64
			for j := 0; j < maxPoints; j++ {
				temp += x(float64(j)) * h(float64(i-j))
			}
			convolution1DResults[i] = temp
		}(i)
	}
	wg.Wait()
	return convolution1DResults
}

func convertToXYs(vect [maxPoints]float64) plotter.XYs {
	pts := make(plotter.XYs, maxPoints)
	for i := 0; i < maxPoints; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(vect[i])
	}
	return pts
}

func Part1() {
	originalSignal := [maxPoints]float64{}
	corruptedSignal := [maxPoints]float64{}
	filter := [maxPoints]float64{}
	for i := 0; i < maxPoints; i++ {
		originalSignal[i] = s(float64(i))
		corruptedSignal[i] = s(float64(i)) + r()
		filter[i] = h(float64(i))
	}

	p := plot.New()
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal"
	err := plotutil.AddLines(p, "sinal original", convertToXYs(originalSignal), "sinal com ruido", convertToXYs(corruptedSignal))
	if err != nil {
		panic(err)
	}
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "imagens-sinais.png"); err != nil {
		panic(err)
	}

	ploterConv := plot.New()
	ploterConv.Title.Text = "Sinal após filtro"
	p.X.Label.Text = "amostras"
	p.Y.Label.Text = "sinal filtrado"
	f := func(f float64) float64 { return s(float64(f)) + r() }
	err = plotutil.AddLines(ploterConv, "Sinal pós filtro", convertToXYs(convolution1D(f, h)))
	if err != nil {
		panic(err)
	}
	if err := ploterConv.Save(4*vg.Inch, 4*vg.Inch, "imagens-sinal-filtrado.png"); err != nil {
		panic(err)
	}
}
